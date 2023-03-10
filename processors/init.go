package processors

import (
	"fmt"

	"github.com/boundedinfinity/docsorter/model"
	"github.com/boundedinfinity/docsorter/util"
	"github.com/oriser/regroup"
)

func (t *ProcessManager) Init(descriptor *model.StatementDescriptor) error {
	for _, line := range descriptor.List {
		if err := util.ValidateLineRegex(*line); err != nil {
			return err
		}
	}

	for _, line := range descriptor.List {
		matcher, err := regroup.Compile(line.Pattern)

		if err != nil {
			return fmt.Errorf("line descriptor error: %v : %w",
				line.Name, err,
			)
		}

		line.Regex = matcher
	}

	return nil
}
