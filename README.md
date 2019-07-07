Create a config file in either /etc/default/sshnotification or $HOME/.sshnotification/sshnotification:

```
{
    "telegram_token": "your-super-secret-bot-token",
    "telegram_id": 123456789
}
```

Add the following line as the last line in /etc/pam.d/sshd:

```
session optional pam_exec.so /path/to/sshloginnotification
```