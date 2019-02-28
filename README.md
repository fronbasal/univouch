# Unifi Voucher Printer
This is a small utility to print Voucher Codes in a fancy fashion.
The Website is designed to print A4 paper. Place icon.png to add a custom icon in front of the voucher.

Quick and dirty(tm).

## Usage
TL;DR: `--serve` enables the webserver, `--generate` generates licenses.
```
usage: univouch --controller=CONTROLLER --username=USERNAME --password=PASSWORD [<flags>]

Flags:
  --help                       Show context-sensitive help (also try --help-long and --help-man).
  --serve                      Serve as printable codes
  --generate                   Generate vouchers
  --controller=CONTROLLER      Controller IP
  --username=USERNAME          Controller username
  --password=PASSWORD          Controller password
  --port=8443                  Controller port
  --vouchers=1                 Amount of vouchers to generate
  --expiration=67              Expiration in minutes
  --note="Generated Vouchers"  Voucher note
```

