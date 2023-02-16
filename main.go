package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"git.fd.io/govpp.git"
	"git.fd.io/govpp.git/api"

	interfaces "mygit.com/myproject/vppbinapi/interface"
	"mygit.com/myproject/vppbinapi/interface_types"
	"mygit.com/myproject/vppbinapi/ip_types"
	sr "mygit.com/myproject/vppbinapi/sr"
	"mygit.com/myproject/vppbinapi/vpe"
)

func GetVPPVersion(ch api.Channel) error {
	fmt.Println("Get VPP Version...")
	request := &vpe.ShowVersion{}
	reply := &vpe.ShowVersionReply{}
	err := ch.SendRequest(request).ReceiveReply(reply)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("VPP Version: %q\n", reply.Version)
	fmt.Println()
	return nil
}

func SrPolicyDump(ch api.Channel) error {
	fmt.Println("Dumping SR Policies")
	time.Sleep(1 * time.Second)

	n := 0
	reqCtx := ch.SendMultiRequest(&sr.SrPoliciesDump{})

	for {
		msg := &sr.SrPoliciesDetails{}
		stop, err := reqCtx.ReceiveReply(msg)
		if stop {
			break
		}
		if err != nil {
			return err
		}
		n++
		fmt.Printf(" - SR Policy #%d: \n", n)
		fmt.Printf("    BSID:      %+v\n", msg.Bsid)
		fmt.Printf("    IsSpray:   %+v\n", msg.IsSpray)
		fmt.Printf("    IsEncap:   %+v\n", msg.IsEncap)
		fmt.Printf("    Fib Table: %+v\n", msg.FibTable)
		fmt.Printf("    SID List:  %+v\n", msg.SidLists[0].Sids)
		//		fmt.Printf("   SID List:  %+v\n", Sids)
	}
	if n == 0 {
		fmt.Println("No Srv6 Policies configured")
	}
	return nil
}

func InterfaceDump(ch api.Channel) error {
	fmt.Println("Dumping interfaces")

	n := 0
	reqCtx := ch.SendMultiRequest(&interfaces.SwInterfaceDump{
		SwIfIndex: ^interface_types.InterfaceIndex(0),
	})
	for {
		msg := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(msg)
		if stop {
			break
		}
		if err != nil {
			return err
		}
		n++
		fmt.Printf(" - interface #%d: %+v", n, msg)
		fmt.Printf(" - interface name: %s", msg.InterfaceName)
		macAddr, err := net.ParseMAC(msg.L2Address.String())
		if err != nil {
			return err
		}
		fmt.Printf(" - interface MAC Address: %s", macAddr)
		fmt.Println()
	}
	return nil
}

func ToVppIP6Address(addr net.IP) ip_types.IP6Address {
	ip := [16]uint8{}
	copy(ip[:], addr)
	return ip
}

//func SrPolicyAdd(ch api.Channel, Bsid ip_types.IP6Address, Isspray bool, Isencap bool, Fibtable int, Sids [16]ip_types.IP6Address, Sidsweight int, Sidslen int) error {
func SrPolicyAdd(ch api.Channel) error {
	fmt.Println("Adding SRv6 Policy")

	IP6BSID := net.ParseIP("2001::3")
	IP6BSIDvpp := ToVppIP6Address(IP6BSID)

	PolicyBsid := ip_types.IP6Address{}
	PolicyBsid = IP6BSIDvpp
	PolicySidList := [16]ip_types.IP6Address{}
	PolicySidList[0] = IP6BSIDvpp
	PolicySidList[1] = IP6BSIDvpp

	request := &sr.SrPolicyAdd{
		BsidAddr: PolicyBsid,
		IsSpray:  false,
		IsEncap:  true,
		FibTable: 0,
		Sids: sr.Srv6SidList{
			NumSids: 2,
			Weight:  1,
			Sids:    PolicySidList,
		},
	}
	response := &sr.SrPolicyAddReply{}
	err := ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Println("SRv6 Policy added!")
	return nil
}

func main() {
	// Connect to VPP
	conn, err := govpp.Connect("/var/run/vpp/vpp-api.sock")
	defer conn.Disconnect()
	if err != nil {
		fmt.Printf("Could not connect: %s\n", err)
		os.Exit(1)
	}

	// Open channel
	ch, err := conn.NewAPIChannel()
	defer ch.Close()
	if err != nil {
		fmt.Printf("Could not open API channel: %s\n", err)
		os.Exit(1)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	err = InterfaceDump(ch)
	if err != nil {
		fmt.Printf("Could not dump interfaces: %s\n", err)
		os.Exit(1)
	}
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println()
		fmt.Println("Please specify your desired action:")
		fmt.Println("If you want to add SRv6 policy, type ADD")
		fmt.Println("If you want to show SRv6 policy, type SHOW")
		fmt.Println("If you want to quit, type QUIT")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		if input == "ADD" {
			fmt.Println("Great! New SRv6 Policy on its way!")
			time.Sleep(1 * time.Second)
			fmt.Println("Please specify the BSID:")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				continue
			}
			input = strings.TrimSuffix(input, "\n")
			policyBSID := ToVppIP6Address(net.ParseIP(input))
			var Isspray bool
			for {
				fmt.Println("Is the policy SPRAY? [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					Isspray = true
					break
				} else if input == "N" {
					Isspray = false
					break
				} else {
					fmt.Println("Please type Y or N")
					continue
				}
			}
			var Isencap bool
			for {
				fmt.Println("Is the policy ENCAP? [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					Isencap = true
					break
				} else if input == "N" {
					Isencap = false
					break
				} else {
					fmt.Println("Please type Y or N")
					continue
				}
			}
			var Fibtable int
			for {
				fmt.Println("Please specify the FIB Table:")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				Fibtable, err = strconv.Atoi(input)
				if err != nil {
					fmt.Printf("Please try again. Cannot convert input to integer: %s\n", err)
					continue
				} else {
					break
				}
			}
			segments := [16]ip_types.IP6Address{}
			fmt.Println("Please insert the SID List [empty input will terminate the List]:")
			i := 0
			n := 1
			for {
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				if len(strings.TrimSpace(input)) == 0 {
					input = strings.TrimSuffix(input, "\n")
					fmt.Println("SID List completed!")
					break
				}
				input = strings.TrimSuffix(input, "\n")
				segments[i] = ToVppIP6Address(net.ParseIP(input))
				i++
				n++
			}
			fmt.Printf(" - SR Policy Ready to be added:\n")
			fmt.Printf("    BSID:      %+v\n", policyBSID)
			fmt.Printf("    IsSpray:   %+v\n", Isspray)
			fmt.Printf("    IsEncap:   %+v\n", Isencap)
			fmt.Printf("    Fib Table: %+v\n", Fibtable)
			fmt.Printf("    SID List:  %+v\n", segments)
			fmt.Println()
			for {
				fmt.Println("Please confirm that this is the policy you want to add: [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					err = SrPolicyAdd(ch)
					//				err = SrPolicyAdd(ch, policyBSID, Isspray, Isencap, Fibtable, segments, Sidsweight, n)
					if err != nil {
						fmt.Printf("Could not add SR Policy: %s\n", err)
						break
					}
					break
				}
				if input == "N" {
					break
				}
				fmt.Println("Please type Y or N")
				continue
			}
		} else if input == "SHOW" {
			err = SrPolicyDump(ch)
			if err != nil {
				fmt.Printf("Could not dump SR Policies: %s\n", err)
				continue
			}
		} else if input == "QUIT" {
			fmt.Println("Exit the session")
			break
		} else {
			fmt.Println("Sorry, type again please")
			continue
		}
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Goodbye!")
	os.Exit(1)
}
