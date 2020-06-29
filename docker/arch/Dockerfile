FROM archlinux

RUN pacman -Sy && \
    pacman -S --noconfirm texinfo texlive-bin \
    	      texlive-core texlive-langcyrillic \
	      texlive-latexextra texlive-pictures \
	      make && \
	      rm -rf /var/cache/pacman/pkg && \
	      rm -rf /var/lib/pacman/sync/*

COPY deps /tex/deps
COPY scripts/bin/* /tex/

RUN /tex/build_deps