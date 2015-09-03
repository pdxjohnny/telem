#include "error.h"

void init_gpgme (gpgme_protocol_t proto)
{
  gpgme_error_t err;

  gpgme_check_version (NULL);
#ifndef HAVE_W32_SYSTEM
  setlocale (LC_ALL, "");
  gpgme_set_locale (NULL, LC_CTYPE, setlocale (LC_CTYPE, NULL));
  gpgme_set_locale (NULL, LC_MESSAGES, setlocale (LC_MESSAGES, NULL));
#endif

  err = gpgme_engine_check_version (proto);
  fail_if_err (err);
}
