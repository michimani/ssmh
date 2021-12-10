ssmh
===

This is a simple CLI tool that outputs the history of connections to Amazon EC2 instances using AWS Session Manager.

# Installation

```bash
brew install michimani/ssmh/ssmh
```

# Usage

```bash
ssmh -h
```

```
                    _
  ___ ___ _ __ ___ | |__
 / __/ __| '_ ' _ \| '_ \
 \__ \__ \ | | | | | | | |
 |___/___/_| |_| |_|_| |_|
 Version: 0.0.0-revision

Usage:
  ssmh [flags] [values]
Flags:
  -n (integer)
    Specifies the number of session histories to retrieve,
    a value greater than 0.

  -r (string)
    Specify the region to retrieve. If not specified, the
    region set in the environment variable "AWS_DEFAULT_REGION"
    will be the target region.

Exapmple:
  ssmh -n 10 -r ap-northeast-1
```

## Output example

```bash
ssmh -r ap-noetheast-1
```

```
+---+---------------------+-----------------+------------------------+----------------+---------------------------+---------------------------+
| # |       TARGET        |  INSTANCE NAME  |       SESSION ID       |     REASON     |        START DATE         |         END DATE          |
+---+---------------------+-----------------+------------------------+----------------+---------------------------+---------------------------+
| 1 | i-01cae5e5a11111111 | test_instance_2 | ssmh-0dbb83f6dc9c0b2ca | 理由は特にない | 2021-12-10T23:32:40+09:00 | 2021-12-10T23:32:49+09:00 |
| 2 | i-05a57d75522222222 | test_instance_1 | ssmh-05f9bf175b72516de |                | 2021-12-10T23:31:59+09:00 | 2021-12-10T23:32:20+09:00 |
| 3 | i-0dbedb80633333333 | test_instance_3 | ssmh-05bbf9a60e1398869 | 焼肉食べたい   | 2021-12-10T23:31:27+09:00 | 2021-12-10T23:31:47+09:00 |
| 4 | i-01cae5e5a11111111 | test_instance_2 | ssmh-01edfc2a84d873870 | 寿司食べたい   | 2021-12-10T23:30:52+09:00 | 2021-12-10T23:31:42+09:00 |
| 5 | i-05a57d75522222222 | test_instance_1 | ssmh-09bff09834feb730f | カレー食べたい | 2021-12-10T23:30:27+09:00 | 2021-12-10T23:31:33+09:00 |
+---+---------------------+-----------------+------------------------+----------------+---------------------------+---------------------------+
```

# Licence

[MIT](https://github.com/michimani/ssmh/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)
