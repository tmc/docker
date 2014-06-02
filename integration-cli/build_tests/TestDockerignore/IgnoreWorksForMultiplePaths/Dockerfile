FROM busybox

ADD . /

# testfile is dockerignored
RUN [ ! -e test_file ]

# test_file2 is not
RUN [ -f test_file2 ]

# dotfiles and dirs can also be hidden
RUN [ ! -e .hidden_file  ]
RUN [ ! -e .hidden_dir  ]
