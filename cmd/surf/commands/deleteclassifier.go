package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/waives/surf/ch360"
	"github.com/waives/surf/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

//go:generate mockery -name "ClassifierDeleter|ClassifierGetter|ClassifierDeleterGetter"

type ClassifierDeleter interface {
	Delete(ctx context.Context, name string) error
}

type ClassifierGetter interface {
	GetAll(ctx context.Context) (ch360.ClassifierList, error)
}

type ClassifierDeleterGetter interface {
	ClassifierDeleter
	ClassifierGetter
}

type DeleteClassifierCmd struct {
	Client         ClassifierDeleterGetter
	ClassifierName string
}

type deleteClassifierArgs struct {
	classifierName string
}

func ConfigureDeleteClassifierCmd(ctx context.Context, deleteCmd *kingpin.CmdClause,
	flags *config.GlobalFlags) {
	args := &deleteClassifierArgs{}
	deleteClassifierCmd := &DeleteClassifierCmd{}

	deleteClassifierCli := deleteCmd.Command("classifier", "Delete waives classifier.").
		Action(func(parseContext *kingpin.ParseContext) error {
			return ExecuteWithMessage(fmt.Sprintf("Deleting classifier '%s'... ", args.classifierName),
				func() error {
					err := deleteClassifierCmd.initFromArgs(args, flags)
					if err != nil {
						return err
					}
					return deleteClassifierCmd.Execute(ctx)
				})
		})

	deleteClassifierCli.
		Arg("name", "The name of the classifier to delete.").
		Required().
		StringVar(&args.classifierName)
}

// Execute runs the 'delete classifier' command.
func (cmd *DeleteClassifierCmd) Execute(ctx context.Context) error {
	classifiers, err := cmd.Client.GetAll(ctx)

	if err != nil {
		return err
	}

	if !classifiers.Contains(cmd.ClassifierName) {
		return errors.New("there is no classifier named '" + cmd.ClassifierName + "'")
	}

	return cmd.Client.Delete(ctx, cmd.ClassifierName)
}

func (cmd *DeleteClassifierCmd) initFromArgs(args *deleteClassifierArgs,
	flags *config.GlobalFlags) error {
	cmd.ClassifierName = args.classifierName

	client, err := initApiClient(flags.ClientId, flags.ClientSecret, flags.LogHttp)

	if err != nil {
		return err
	}

	cmd.Client = client.Classifiers
	return nil
}
