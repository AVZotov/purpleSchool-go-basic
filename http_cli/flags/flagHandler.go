package flags

import (
	"errors"
	"flag"
	"fmt"
)

func httpRequest() []string {
	return []string{"create", "get", "update", "delete"}
}

type Flags struct {
	List     bool
	Method   string
	Filepath string
	BinName  string
	ID       string
}

func (flags *Flags) ReadAll() bool {
	return flags.List
}

func (flags *Flags) Request() string {
	return flags.Method
}

func (flags *Flags) Path() string {
	return flags.Filepath
}

func (flags *Flags) Name() string {
	return flags.BinName
}

func (flags *Flags) Id() string {
	return flags.ID
}

func GetFlags() (*Flags, error) {
	flags := flagParser()
	var err error
	httpMethodFunc := map[string]func(*Flags) error{
		httpRequest()[0]: validateCreateFlag,
		httpRequest()[1]: validateGetAndDeleteFlag,
		httpRequest()[2]: validateUpdateFlag,
		httpRequest()[3]: validateGetAndDeleteFlag,
	}

	if flags.List {
		if err = validateListFlag(flags); err != nil {
			return nil, err
		}
		return flags, nil
	}
	if err = httpMethodFunc[flags.Method](flags); err != nil {
		return nil, err
	}
	return flags, nil
}

func validateListFlag(flags *Flags) error {
	if flags.Method != "" || flags.Filepath != "" || flags.BinName != "" || flags.ID != "" {
		return errors.New("too many arguments with --list flag")
	}
	return nil
}

func validateGetAndDeleteFlag(flags *Flags) error {
	if flags.ID == "" {
		return errors.New("missing id flag")
	}
	if flags.Filepath != "" && flags.BinName != "" {
		return fmt.Errorf("too many arguments with --%s flag", flags.Method)
	}
	return nil
}

func validateUpdateFlag(flags *Flags) error {
	if flags.ID == "" {
		return errors.New("missing id flag")
	}
	if flags.Filepath == "" {
		return errors.New("missing file flag")
	}
	if flags.BinName != "" {
		return fmt.Errorf("too many arguments with --%s flag", flags.Method)
	}
	return nil
}

func validateCreateFlag(flags *Flags) error {
	if flags.Filepath == "" {
		return errors.New("missing file flag")
	}
	if flags.BinName == "" {
		return errors.New("missing name flag")
	}
	if flags.ID != "" {
		return fmt.Errorf("too many arguments with --%s flag", flags.Method)
	}
	return nil
}

func flagParser() *Flags {
	flags := Flags{}

	flag.BoolVar(&flags.List, "list", false, "lists all bins from the local storage")
	for _, method := range httpRequest() {
		flag.BoolFunc(method, fmt.Sprintf("http %q request. Must be only one flag of that type", method),
			func(s string) error {
				if flags.Method != "" {
					return fmt.Errorf(
						"only one flag is allowed with supplied --%q & --%q", flags.Method, method)
				}
				if s == "true" {
					flags.Method = method
				}
				return nil
			})
	}
	flag.StringVar(&flags.Filepath, "file", "", "path to file with bin stored")
	flag.StringVar(&flags.BinName, "name", "", "name of the bin")
	flag.StringVar(&flags.ID, "id", "", "id of the bin")

	flag.Parse()
	return &flags
}
