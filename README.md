# GoVPP for SRv6

Uses GO client interacting with VPP's API to provision SRv6 policies and Steer Traffic.

#### Quick Demo (4 min):

https://cisco.webex.com/cisco/ldr.php?RCID=32dd1b9b654e2755b8255430e53bbbd0

Pswd: yT7PSs9a

#### Prerequisites:
- GoVPP: https://wiki.fd.io/view/GoVPP
- Golang: https://go.dev/
- VPP: https://wiki.fd.io/view/VPP

In order to import the code to your personale GitHub repository:
1. Clone the Repository into your personal GitHub
2. Open the code in code editing application (ex. Visual Studio)
3. Make your code changes
4. Find and Replace all the achiarato/GoVPP links with your GitHub repository path
5. Git add, Git commit, Git push

In order to use the code once downloaded:
```
1. cd GoVPP/
2. go build
3. sudo ./GoVPP 
```

GoVPP will generate interactive CLI output.:
```
GoVPP Ready to Rock!

Please specify your desired action:
If you want to get VPP detailes, type DET
If you want to add SRv6 policy, type ADD
If you want to add SR Steer policy, type STEER
If you want to show SRv6 policy, type SHOW
If you want to quit, type QUIT
```

DET: Provide details for VPP Version and VPP Interfaces

ADD: Define and Provision SRv6 Policy

STEER: Define and Steer Traffic to existing SRv6 Policy

SHOW: Show existing SRv6 Policy

QUIT: Close the GoVPP session
