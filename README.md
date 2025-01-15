# GroupMe Viewer

Displays exported GroupMe chats.

### Running

1. Execute `groupme-viewer` (Linux) or `groupme-viewer.exe` (Windows).
2. Go to `localhost:3000` in your browser.

Building from source is left as an exercise for the reader.

### Adding a Group Chat

1. Go to the [GroupMe website](https://groupme.com) in your browser (with a computer).
2. Click your profile picture in the bottom left.
3. Click "Profile" if it's not already selected.
4. Scroll to the bottom and click "Export my data".
5. Click "Create Export".
6. Select "Message Data", "Attachments", and "Attachments From Others"; click "Next".
7. Select any groups you want to export.
    1. For very large/old groups, I recommend selecting only one at a time (then going through these steps again for other chats after you've downloaded it) as GroupMe exports will fail if they're too large/take too long.
    2. If you see checkboxes but not group names, press Ctrl+A to highlight all so they're visible - GroupMe's formatting is messed up.
8. Click "Export".
9. Click "Back" (there's a message here saying they'll email you - this is a lie).
10. Under "Manage Export", if there are file downloads available, proceed.  Otherwise, click the refresh button in the top right of the section until there are (may take many minutes).
11. Download all the files (small chats may have only one).
12. Extract the files to `./data`. You should have a directory structure like this:
```
./data
├── 123456789
│   ├── conversation.json
│   ├── gallery
│   │   ├── message_id_1280.123456789.jpg
│   │   ├── message_id_640.123456789.jpg
│   └── message.json
└── 987654321
    ├── conversation.json
    ├── gallery
    │   ├── message_id_1280.123456789.jpg
    │   ├── message_id_640.123456789.jpg
    └── message.json
```
13. If you know what you're doing, you can use the `scrape_attachments.ipynb` notebook to recover some additional images from GroupMe's servers (it gets you about another year). I take no responsibility if you get banned for downloading a million images.
14. Repeat for any other chats you want to add.

### Missing Features

No support for or plans to support:

* Polls
* Uploaded Files
* Avatars
* Group Images and Custom Reaction Emoji
* Probably other stuff

Partial Support:

* Replies (link only - no preview)

Other Limitations:

* Entire chats are loaded into memory. This has worked fine for me (~100k messages), but you may run into issues.
