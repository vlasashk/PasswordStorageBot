package configs

const (
	StartMsg string = `Bot is active and running!
Use /help command to familiarize with functionality`
	SetMsg  string = "Enter service name to add credentials"
	GetMsg  string = "Choose service name to edit credentials"
	DelMsg  string = "Choose service name to remove credentials"
	HelpMsg string = `/start - used to initialize the menu.
/set - used to add new service with new credentials. 
Each user is allowed to have 20 services in total.
Each service is allowed to have 1 set of credentials (login : password).
/get - used to get the credentials of the service you have already added.
/del - used to remove the service and it's credentials from database.`
	DefaultMsg string = "Unknown command. Try /get, /set or /del"
	LoginMsg   string = "Enter Login for "
	PassMsg    string = "Enter Password for "
	SuccessMsg string = "Credentials were successfully added for "
	LimitMsg   string = `You are out of service limit. 
Currently it is allowed to have no more than 20 services.
In order to add new service you must remove one`
	ExistMsg string = ` service already exist in storage.
If you want to change credentials for this service you must first remove it and then add it once more`
	EmptyMsg string = `Your account doesn't have any added services.
Try /set to add new one`
	MissingMsg string = `Your account doesn't have any available credentials for `
	LenErrMsg  string = `Your input is too long. Input must be no longer then 50 characters`
)
