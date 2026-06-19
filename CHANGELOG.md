Changelog
=========

v0.5.0
------
Jun 19, 2026

- Improve handling of non-UTF-8 input (#14)
  - Attempt to decode invalid UTF-8 using the current system encoding
  - If decoding fails, preserve printable ASCII (`0x20`–`0x7E`) and
    display all other bytes as `\xNN`

v0.4.0
------
Jun 14, 2026

- Add options: `-strip-cr` and `-show-control` (#11)
- Display U+2400 - U+241F symbols for control characters (default) (#12)

v0.3.2
------
May 10, 2026

- EventLoop methods: Open/Close TTY only when not already opened (#8)
- Use "go-ttyadapter/fav" instead of "go-ttyadapter/tty8pe" (#9)

v0.3.1
------
Apr 30, 2026

- Application now shutdowns immediately when no files are specified and standard input is not redirected (#6)

v0.3.0
------
Apr 28, 2026

- Add Ctrl-L to redraw the screen (#4)
- Add Session.ClearCache to force re-rendering (#4)

v0.2.0
-------
Apr 27, 2026

- Fix: nemo CLI ignored the first entered key (#2)
- Update "mattn/go-tty" to v2.0.0
- Update "nyaosorg/go-ttyadapter" to v0.6.0
- Update "nyaosorg/go-readline-ny" to v1.15.0

v0.1.0
------
Apr 26, 2026

- Initial version
