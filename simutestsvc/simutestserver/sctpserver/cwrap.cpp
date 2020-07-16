#include "cwrap.h"
#include "sctpserver.h"

void startSctp() {
    SctpServer server;
    server.start();
}