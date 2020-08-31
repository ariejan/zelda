# Zelda

This is a cli utility for Hertek Connect Link.

# Status

This is a work-in-progress, thing may or may not work. Do not use for production purposes.

# Usage

This should work on any Linux/BSD/macos system a recent Go version. It is developed using Go 1.15.

To setup your credentials for Connect Link, you'll need to make a config file at `~/.config/zelda.yaml` 
with the following content:

```yaml
zelda:
  server_url: https://api.example.com
  username: <connect-link-integrator-username>
  password: <connect-link-integrator-password>
```

Make sure to make this file readable only by you with `chmod 600 ~/.config/zelda.yaml`.

When you have your configuration setup, you can test your setup with: `zelda ping`.

# Disclaimer

This Hertek Connect Link client is not developed nor endorsed by Hertek. Use at your own risk. 

# Contributors

 * [Ariejan de Vroom](https://www.devroom.io)

# License

MIT License

Copyright (c) 2020 Ariejan de Vroom

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
