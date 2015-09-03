#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <gpgme.h>

#include "gpg_opts.h"
#include "error.h"
#include "debug.h"
#include "init.h"

#ifndef PGM
#define PGM "import"
#endif

int main (int argc, char **argv)
{
  gpgme_error_t err;
  gpgme_ctx_t context;
  telem_gpg_opts options;
  int url_mode = 0;
  int nul_mode = 0;
  gpgme_import_result_t impres;
  gpgme_data_t data;

  setup(&context, &options);

  argv++;
  for (; argc > 1; --argc, argv++)
    {
      printf ("reading file `%s'\n", *argv);
      err = gpgme_data_new_from_file(&data, *argv, 1);
      fail_if_err(err);

      if (url_mode)
        gpgme_data_set_encoding(
          data,
          (nul_mode? GPGME_DATA_ENCODING_URL0 : GPGME_DATA_ENCODING_URL)
        );

      err = gpgme_op_import(context, data);
      fail_if_err(err);
      impres = gpgme_op_import_result(context);
      if (!impres)
        {
          fprintf(stderr, PGM ": no import result returned\n");
          exit(1);
        }
      print_import_result(impres);

      gpgme_data_release(data);
    }

  gpgme_release(context);
  return 0;
}
