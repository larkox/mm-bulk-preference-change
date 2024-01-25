# mm-bulk-preference-change

mm-bulk-preference-change is a tool to update the preferences of Mattermost users in bulk.

## Requirements
You need mmctl. If it is not installed on the PATH, you can provide a path to the mmctl executable.
Also, mmctl must be authenticated. See the `mmctl auth -h` for more information.

## Usage
The command will read the user emails from a file. This file must have the emails separated by white spaces, tabs or new lines.

You can use the command as:
```
mm-bulk-preference-change emails.txt favorite_channel ungsg3mx77gkdxqnacbo1aa6de true
```

If you need to define the path to the mmctl executable, you can use the `--mmctlPath` flag. Example:
```
mm-bulk-preference-change --mmctlPath ./mattermost/server/bin/mmctl emails.txt favorite_channel ungsg3mx77gkdxqnacbo1aa6de true
```