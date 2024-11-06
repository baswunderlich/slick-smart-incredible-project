# README

## About

This is the official Wails Svelte template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Site for key pair generation

https://8gwifi.org/RSAFunctionality?rsasignverifyfunctions=rsasignverifyfunctions&keysize=2048

## VCs 
Currently, there are four VCs stored in monsi. 

- bachelorDegree.json
- examVC.json
- profVC.json
- studentVC.json

The only valid one is the studentVC.json representing a studenConfirmation

The others have the following problems:

- bachelorDegree.json   (too old)
- examVC.json           (invalid signature)
- profVC.json           (too old and invalid signature)