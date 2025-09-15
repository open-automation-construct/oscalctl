# stigctl
A Golang STIG Automation Tool

##

Idea's written down

Take a config value to point to a .cklb file (Created via STIGVIEWER), this could live in the git repo with whatever application.

Write integration with OPA, Kyverno etc. to write automation to auto-check and fill out the checklist. 

Eventually will need a way to fetch live configs (API, kubernetes, file etc.)