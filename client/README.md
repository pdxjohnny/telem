Telem Client
---

This is where client apis are stored

```bash
gcc *.c \
  `gpgme-config --libs` \
  `gpg-error-config --libs` \
  `libassuan-config --libs` \
  --static -o client
file ./client
```
