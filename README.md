# go-vpp-sr

Uses VPP's go-vpp API to provision SRv6 policies

Tested with golang 1.19.5

1. go build

```
cd go-vpp-sr
go build -buildvcs=false
```

2. usage:

```
sudo ./go-vpp-sr 
```

3. Interactive ouput
```
Dumping interfaces
 - interface #1: &{SwIfIndex:0 SupSwIfIndex:0 L2Address:00:00:00:00:00:00 Flags:IfStatusFlags(0) Type:IF_API_TYPE_HARDWARE LinkDuplex:LINK_DUPLEX_API_UNKNOWN LinkSpeed:0 LinkMtu:0 Mtu:[0 0 0 0] SubID:0 SubNumberOfTags:0 SubOuterVlanID:0 SubInnerVlanID:0 SubIfFlags:SubIfFlags(0) VtrOp:0 VtrPushDot1q:0 VtrTag1:0 VtrTag2:0 OuterTag:0 BDmac:00:00:00:00:00:00 BSmac:00:00:00:00:00:00 BVlanid:0 ISid:0 InterfaceName:local0 InterfaceDevType:local Tag:} - interface name: local0 - interface MAC Address: 00:00:00:00:00:00
 - interface #2: &{SwIfIndex:1 SupSwIfIndex:1 L2Address:52:54:68:45:a5:68 Flags:IF_STATUS_API_FLAG_ADMIN_UP|IF_STATUS_API_FLAG_LINK_UP Type:IF_API_TYPE_HARDWARE LinkDuplex:LINK_DUPLEX_API_FULL LinkSpeed:4294967295 LinkMtu:9000 Mtu:[9000 0 0 0] SubID:0 SubNumberOfTags:0 SubOuterVlanID:0 SubInnerVlanID:0 SubIfFlags:SubIfFlags(0) VtrOp:0 VtrPushDot1q:0 VtrTag1:0 VtrTag2:0 OuterTag:0 BDmac:00:00:00:00:00:00 BSmac:00:00:00:00:00:00 BVlanid:0 ISid:0 InterfaceName:GigabitEthernet0/7/0 InterfaceDevType:dpdk Tag:} - interface name: GigabitEthernet0/7/0 - interface MAC Address: 52:54:68:45:a5:68
 - interface #3: &{SwIfIndex:2 SupSwIfIndex:2 L2Address:00:00:00:00:00:00 Flags:IF_STATUS_API_FLAG_ADMIN_UP|IF_STATUS_API_FLAG_LINK_UP Type:IF_API_TYPE_HARDWARE LinkDuplex:LINK_DUPLEX_API_UNKNOWN LinkSpeed:0 LinkMtu:9000 Mtu:[9000 0 0 0] SubID:0 SubNumberOfTags:0 SubOuterVlanID:0 SubInnerVlanID:0 SubIfFlags:SubIfFlags(0) VtrOp:0 VtrPushDot1q:0 VtrTag1:0 VtrTag2:0 OuterTag:0 BDmac:00:00:00:00:00:00 BSmac:00:00:00:00:00:00 BVlanid:0 ISid:0 InterfaceName:host-vpp-in InterfaceDevType:af-packet Tag:} - interface name: host-vpp-in - interface MAC Address: 00:00:00:00:00:00

Please specify your desired action:
If you want to add SRv6 policy, type ADD
If you want to show SRv6 policy, type SHOW
If you want to quit, type QUIT
ADD
Great! New SRv6 Policy on its way!
Please specify the BSID:
1::1
ToVppIP6Address [0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
BSID:  1::1
Is the policy SPRAY? [Y/N]
N
Is the policy ENCAP? [Y/N]
Y
Please specify the FIB Table:
0
Please insert the SID List [empty input will terminate the List]:
fc00:0:41:5:13:61:66::
ToVppIP6Address [252 0 0 0 0 65 0 5 0 19 0 97 0 102 0 0]

SID List completed!
 - SR Policy Ready to be added:
    BSID:      1::1
    IsSpray:   false
    IsEncap:   true
    Fib Table: 0
    SID List:  [fc00:0:41:5:13:61:66:0 :: :: :: :: :: :: :: :: :: :: :: :: :: :: ::]

Please confirm that this is the policy you want to add: [Y/N]
Y
Adding SRv6 Policy
SID0 fc00:0:41:5:13:61:66:0
SID1 ::
SRv6 Policy added!

Please specify your desired action:
If you want to add SRv6 policy, type ADD
If you want to show SRv6 policy, type SHOW
If you want to quit, type QUIT
^Z
[5]+  Stopped                 sudo ./go-vpp-sr
```

4. vppctl to display policies:
```
show sr policies 
```

5. vppctl to delete policies:
```
sr policy del bsid 1::1
```