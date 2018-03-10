package main

// Code generated but intended to be editted by the user

// ActionCustom lets you define a set of cutom methods and fields on action, which won't get overwritten!
type ActionCustom struct {
	// Define fields that come from third party libs, etc, here
}

// Run is the custom run method to execute the Action that you design!
func (a *Action) Run(input *ActionInput) (*ActionOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	// Put any custom code you want in here!
	return &ActionOutput{}, nil
}

// Validate will return an error if the input struct isn't filled in correctly.
func (i *ActionInput) Validate() error {
	// Allow the user to write custom validation code for their inputs!
	return nil
}
