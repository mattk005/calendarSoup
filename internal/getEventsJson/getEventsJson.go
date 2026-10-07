package getEventsJson

import (
	"os"

	"wfm-json-to-ics/internal/config"
)

// probably break this into "update single user, update all users", seems like alot of work for something that wont run very often
func GetEventsJSON(args []string, cfg *config.Config) (string, error) {
	if len(args) == 0 {
		// this is just left over from when this was a CLI so it doesn't make a tonne of sence
		return "this app requires arguments:\n\t-updateUsers\n\t\tWill update the database users table from users.json\notherwise, this app requires a JSESSIONID and ROUTEID token from the app.", nil
	}

	switch args[0] {
	case "-updateUsers":
		{
			if err := cfg.GetUserData("users.json"); err != nil {
				return "", err
			}
			os.Exit(0)
		}
	case "-printUsers":
		{
			resp, err := cfg.PrintUsers()
			return resp, err
		}

	}

	var err error
	cfg.Tokens, err = config.GetTokens(args)
	if err != nil {
		return "", err
	}

	if err := updateFromWFM(cfg); err != nil {
		return "", err
	}
	return "Updated Users", nil
}
