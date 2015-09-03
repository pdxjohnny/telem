// Options for telem_gpg
#ifndef TELEM_OPT
#define TELEM_OPT 1

#define TELEM_STRING_SIZE 200

// Defaults for linux
  #define TELEM_OPT_KEYRING_DIR "../../keys/client/keyring_dir"


typedef struct telem_gpg_opts_struct {
  char keyring_dir[TELEM_STRING_SIZE];
} telem_gpg_opts;

#endif
