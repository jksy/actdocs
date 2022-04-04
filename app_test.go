package actdocs

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAppRunWithGenerate(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{"generate", "--sort", "testdata/valid-workflow.yml"},
			expected: expectedGenerateWithSortWorkflow,
		},
		{
			args:     []string{"generate", "--sort-by-name", "testdata/valid-workflow.yml"},
			expected: expectedGenerateWithSortByNameWorkflow,
		},
		{
			args:     []string{"generate", "--sort", "--omit", "testdata/valid-empty-workflow.yml"},
			expected: expectedGenerateWithOmitWorkflow,
		},
		{
			args:     []string{"generate", "--sort", "testdata/valid-empty-workflow.yml"},
			expected: expectedGenerateWithEmptyWorkflow,
		},
		{
			args:     []string{"generate", "--sort", "--format=json", "testdata/valid-workflow.yml"},
			expected: expectedGenerateWithSortFormatJsonWorkflow,
		},
		{
			args:     []string{"generate", "--sort", "--format=json", "testdata/valid-empty-workflow.yml"},
			expected: expectedGenerateWithEmptyFormatJsonWorkflow,
		},
		{
			args:     []string{"generate", "--format=json", "testdata/valid-empty-workflow.yml"},
			expected: expectedGenerateWithEmptyFormatJsonWorkflow,
		},
		{
			args:     []string{"generate", "--sort", "testdata/valid-action.yml"},
			expected: expectedGenerateWithSortAction,
		},
		{
			args:     []string{"generate", "--sort-by-name", "testdata/valid-action.yml"},
			expected: expectedGenerateWithSortByNameAction,
		},
		{
			args:     []string{"generate", "--sort", "--omit", "testdata/valid-empty-action.yml"},
			expected: expectedGenerateWithOmitAction,
		},
		{
			args:     []string{"generate", "--sort", "testdata/valid-empty-action.yml"},
			expected: expectedGenerateWithEmptyAction,
		},
		{
			args:     []string{"generate", "--sort", "--format=json", "testdata/valid-action.yml"},
			expected: expectedGenerateWithSortFormatJsonAction,
		},
		{
			args:     []string{"generate", "--sort", "--format=json", "testdata/valid-empty-action.yml"},
			expected: expectedGenerateWithEmptyFormatJsonAction,
		},
		{
			args:     []string{"generate", "--format=json", "testdata/valid-empty-action.yml"},
			expected: expectedGenerateWithEmptyFormatJsonAction,
		},
	}

	app := NewApp("test", "", "", "")
	for _, tc := range cases {
		outWriter := &bytes.Buffer{}
		inOut := NewIO(os.Stdin, outWriter, os.Stderr)
		err := app.Run(tc.args, inOut.InReader, inOut.OutWriter, inOut.ErrWriter)

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", strings.Join(tc.args, " "), err)
		}

		if diff := cmp.Diff(outWriter.String(), tc.expected); diff != "" {
			t.Errorf("%s: unexpected out: \n%s", strings.Join(tc.args, " "), diff)
		}
	}
}

const expectedGenerateWithSortWorkflow = `## Inputs

| Name | Description | Type | Default | Required |
| :--- | :---------- | :--- | :------ | :------: |
| full-string | The full string value. | ` + "`string`" + ` | ` + "``" + ` | yes |
| required-and-description | The required and description value. | n/a | n/a | yes |
| default-and-type |  | ` + "`string`" + ` | ` + "`foo`" + ` | no |
| empty |  | n/a | n/a | no |
| full-boolean | The full boolean value. | ` + "`boolean`" + ` | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`number`" + ` | ` + "`5`" + ` | no |

## Secrets

| Name | Description | Required |
| :--- | :---------- | :------: |
| alternative-required-secret | The alternative required secret value. | yes |
| required-secret | The required secret value. | yes |
| empty |  | no |
| not-required-secret | The not required secret value. | no |
| without-required-secret | The not required secret value. | no |
`

const expectedGenerateWithOmitWorkflow = "\n"

const expectedGenerateWithEmptyWorkflow = `## Inputs

N/A

## Secrets

N/A
`

const expectedGenerateWithSortByNameWorkflow = `## Inputs

| Name | Description | Type | Default | Required |
| :--- | :---------- | :--- | :------ | :------: |
| default-and-type |  | ` + "`string`" + ` | ` + "`foo`" + ` | no |
| empty |  | n/a | n/a | no |
| full-boolean | The full boolean value. | ` + "`boolean`" + ` | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`number`" + ` | ` + "`5`" + ` | no |
| full-string | The full string value. | ` + "`string`" + ` | ` + "``" + ` | yes |
| required-and-description | The required and description value. | n/a | n/a | yes |

## Secrets

| Name | Description | Required |
| :--- | :---------- | :------: |
| alternative-required-secret | The alternative required secret value. | yes |
| empty |  | no |
| not-required-secret | The not required secret value. | no |
| required-secret | The required secret value. | yes |
| without-required-secret | The not required secret value. | no |
`

const expectedGenerateWithSortFormatJsonWorkflow = `{
  "inputs": [
    {
      "name": "full-string",
      "default": "",
      "description": "The full string value.",
      "required": "true",
      "type": "string"
    },
    {
      "name": "required-and-description",
      "default": null,
      "description": "The required and description value.",
      "required": "true",
      "type": null
    },
    {
      "name": "default-and-type",
      "default": "foo",
      "description": null,
      "required": null,
      "type": "string"
    },
    {
      "name": "empty",
      "default": null,
      "description": null,
      "required": null,
      "type": null
    },
    {
      "name": "full-boolean",
      "default": "true",
      "description": "The full boolean value.",
      "required": "false",
      "type": "boolean"
    },
    {
      "name": "full-number",
      "default": "5",
      "description": "The full number value.",
      "required": "false",
      "type": "number"
    }
  ],
  "secrets": [
    {
      "name": "alternative-required-secret",
      "description": "The alternative required secret value.",
      "required": "true"
    },
    {
      "name": "required-secret",
      "description": "The required secret value.",
      "required": "true"
    },
    {
      "name": "empty",
      "description": null,
      "required": null
    },
    {
      "name": "not-required-secret",
      "description": "The not required secret value.",
      "required": "false"
    },
    {
      "name": "without-required-secret",
      "description": "The not required secret value.",
      "required": null
    }
  ]
}
`

const expectedGenerateWithEmptyFormatJsonWorkflow = `{
  "inputs": [],
  "secrets": []
}
`

const expectedGenerateWithSortAction = `## Inputs

| Name | Description | Default | Required |
| :--- | :---------- | :------ | :------: |
| full-string | The full string value. | ` + "`Default value`" + ` | yes |
| description-only | The description without default and required. | n/a | no |
| empty |  | n/a | no |
| full-boolean | The full boolean value. | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`5`" + ` | no |

## Outputs

| Name | Description |
| :--- | :---------- |
| only-value |  |
| with-description | The output value with description. |
`

const expectedGenerateWithSortByNameAction = `## Inputs

| Name | Description | Default | Required |
| :--- | :---------- | :------ | :------: |
| description-only | The description without default and required. | n/a | no |
| empty |  | n/a | no |
| full-boolean | The full boolean value. | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`5`" + ` | no |
| full-string | The full string value. | ` + "`Default value`" + ` | yes |

## Outputs

| Name | Description |
| :--- | :---------- |
| only-value |  |
| with-description | The output value with description. |
`

const expectedGenerateWithOmitAction = "\n"

const expectedGenerateWithEmptyAction = `## Inputs

N/A

## Outputs

N/A
`

const expectedGenerateWithSortFormatJsonAction = `{
  "description": "This is a test Custom Action for actdocs.",
  "inputs": [
    {
      "Name": "full-string",
      "Default": "Default value",
      "Description": "The full string value.",
      "Required": "true"
    },
    {
      "Name": "description-only",
      "Default": null,
      "Description": "The description without default and required.",
      "Required": null
    },
    {
      "Name": "empty",
      "Default": null,
      "Description": null,
      "Required": null
    },
    {
      "Name": "full-boolean",
      "Default": "true",
      "Description": "The full boolean value.",
      "Required": "false"
    },
    {
      "Name": "full-number",
      "Default": "5",
      "Description": "The full number value.",
      "Required": "false"
    }
  ],
  "outputs": [
    {
      "Name": "only-value",
      "Description": null
    },
    {
      "Name": "with-description",
      "Description": "The output value with description."
    }
  ]
}
`

const expectedGenerateWithEmptyFormatJsonAction = `{
  "description": null,
  "inputs": [],
  "outputs": []
}
`

func TestAppRunWithInject(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{"inject", "--sort", "--dry-run", "--file=testdata/output.md", "testdata/valid-workflow.yml"},
			expected: expectedInjectWithSortWorkflow,
		},
		{
			args:     []string{"inject", "--sort", "--dry-run", "--file=testdata/output.md", "testdata/valid-empty-workflow.yml"},
			expected: expectedInjectWithEmptyWorkflow,
		},
		{
			args:     []string{"inject", "--sort", "--dry-run", "--omit", "--file=testdata/output.md", "testdata/valid-empty-workflow.yml"},
			expected: expectedInjectWithOmitWorkflow,
		},
		{
			args:     []string{"inject", "--sort", "--dry-run", "--file=testdata/output.md", "testdata/valid-action.yml"},
			expected: expectedInjectWithSortAction,
		},
		{
			args:     []string{"inject", "--sort", "--dry-run", "--file=testdata/output.md", "testdata/valid-empty-action.yml"},
			expected: expectedInjectWithEmptyAction,
		},
		{
			args:     []string{"inject", "--sort", "--dry-run", "--omit", "--file=testdata/output.md", "testdata/valid-empty-action.yml"},
			expected: expectedInjectWithOmitAction,
		},
	}

	app := NewApp("test", "", "", "")
	for _, tc := range cases {
		outWriter := &bytes.Buffer{}
		inOut := NewIO(os.Stdin, outWriter, os.Stderr)
		err := app.Run(tc.args, inOut.InReader, inOut.OutWriter, inOut.ErrWriter)

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", strings.Join(tc.args, " "), err)
		}

		if diff := cmp.Diff(outWriter.String(), tc.expected); diff != "" {
			t.Errorf("%s: unexpected out: \n%s", strings.Join(tc.args, " "), diff)
		}
	}
}

const expectedInjectWithSortWorkflow = `# Output test

## Header

This is a header.

<!-- actdocs start -->

## Inputs

| Name | Description | Type | Default | Required |
| :--- | :---------- | :--- | :------ | :------: |
| full-string | The full string value. | ` + "`string`" + ` | ` + "``" + ` | yes |
| required-and-description | The required and description value. | n/a | n/a | yes |
| default-and-type |  | ` + "`string`" + ` | ` + "`foo`" + ` | no |
| empty |  | n/a | n/a | no |
| full-boolean | The full boolean value. | ` + "`boolean`" + ` | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`number`" + ` | ` + "`5`" + ` | no |

## Secrets

| Name | Description | Required |
| :--- | :---------- | :------: |
| alternative-required-secret | The alternative required secret value. | yes |
| required-secret | The required secret value. | yes |
| empty |  | no |
| not-required-secret | The not required secret value. | no |
| without-required-secret | The not required secret value. | no |

<!-- actdocs end -->

## Footer

This is a footer.
`

const expectedInjectWithEmptyWorkflow = `# Output test

## Header

This is a header.

<!-- actdocs start -->

## Inputs

N/A

## Secrets

N/A

<!-- actdocs end -->

## Footer

This is a footer.
`

const expectedInjectWithOmitWorkflow = `# Output test

## Header

This is a header.

<!-- actdocs start -->
<!-- actdocs end -->

## Footer

This is a footer.
`

const expectedInjectWithSortAction = `# Output test

## Header

This is a header.

<!-- actdocs start -->

## Inputs

| Name | Description | Default | Required |
| :--- | :---------- | :------ | :------: |
| full-string | The full string value. | ` + "`Default value`" + ` | yes |
| description-only | The description without default and required. | n/a | no |
| empty |  | n/a | no |
| full-boolean | The full boolean value. | ` + "`true`" + ` | no |
| full-number | The full number value. | ` + "`5`" + ` | no |

## Outputs

| Name | Description |
| :--- | :---------- |
| only-value |  |
| with-description | The output value with description. |

<!-- actdocs end -->

## Footer

This is a footer.
`

const expectedInjectWithEmptyAction = `# Output test

## Header

This is a header.

<!-- actdocs start -->

## Inputs

N/A

## Outputs

N/A

<!-- actdocs end -->

## Footer

This is a footer.
`

const expectedInjectWithOmitAction = `# Output test

## Header

This is a header.

<!-- actdocs start -->
<!-- actdocs end -->

## Footer

This is a footer.
`
