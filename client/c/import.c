#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <gpgme.h>

#include "error.h"
#include "debug.h"

#ifndef PGM
#define PGM "import"
#endif

int main (int argc, char **argv)
{
  gpgme_error_t err;
  gpgme_ctx_t ctx;
  int url_mode = 0;
  int nul_mode = 0;
  gpgme_import_result_t impres;
  gpgme_data_t data;

  init_gpgme (GPGME_PROTOCOL_OpenPGP);

  err = gpgme_new (&ctx);
  fail_if_err (err);
  gpgme_set_protocol (ctx, GPGME_PROTOCOL_OpenPGP);

  argv++;
  for (; argc > 1; --argc, argv++)
    {
      printf ("reading file `%s'\n", *argv);
      err = gpgme_data_new_from_file (&data, *argv, 1);
      fail_if_err (err);

      if (url_mode)
        gpgme_data_set_encoding (data, (nul_mode? GPGME_DATA_ENCODING_URL0
                                        : GPGME_DATA_ENCODING_URL));

      err = gpgme_op_import (ctx, data);
      fail_if_err (err);
      impres = gpgme_op_import_result (ctx);
      if (!impres)
        {
          fprintf (stderr, PGM ": no import result returned\n");
          exit (1);
        }
      print_import_result (impres);

      gpgme_data_release (data);
    }

  gpgme_release (ctx);
  return 0;
}
