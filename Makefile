CMARKDIR=$(shell pwd)/cmark-gfm
CMDPREFIX=$(shell scripts/prefix.sh)
INSTALL_PREFIX=$(shell pwd)/.root

.PHONY: cmark clean test


clean:
	rm -rf $(INSTALL_PREFIX)
	cd $(CMARKDIR) && make distclean


test: cmark
	export CGO_ENABLED=0 && \
	export PATH=$(INSTALL_PREFIX)/bin:$(PATH) && \
	$(CMDPREFIX) go test -cover -v ./...


# ------------------------------------------------------------------------------


cmark: $(INSTALL_PREFIX)/bin/cmark-gfm


$(INSTALL_PREFIX)/bin/cmark-gfm:
	mkdir -p $(CMARKDIR)/build && \
	cd $(CMARKDIR)/build && \
	$(CMDPREFIX) cmake .. \
		-DCMAKE_BUILD_TYPE=Release \
		-DCMAKE_INSTALL_PREFIX=$(INSTALL_PREFIX) \
		-DCMAKE_EXPORT_COMPILE_COMMANDS=ON \
		-DCMARK_SHARED=OFF

	cd $(CMARKDIR) && \
	$(CMDPREFIX) make cmake_build install
