# This file is only used for running automated rpm build tests for PRs
# The canonical dist-git source can be found at:
# https://src.fedoraproject.org/rpms/cri-o/blob/master/f/cri-o.spec

%if 0%{?with_debug}
%global _dwz_low_mem_die_limit 0
%else
%global debug_package   %{nil}
%endif

%global provider        github
%global provider_tld    com
%global project         kubernetes-incubator
%global repo            cri-o
%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}
%global service_name crio

Name:           cri-o-testonly
Version:        1.8.0
Release:        1.dev.testonly%{?dist}
Summary:        OCI-based implementation of Kubernetes Container Runtime Interface
License:        ASL 2.0
URL:            https://%{provider_prefix}
Source0:        %{name}_%{version}.orig.tar.gz
ExcludeArch:    ppc64
# If go_compiler is not set to 1, there is no virtual provide. Use golang instead.
BuildRequires:  %{?go_compiler:compiler(go-compiler)}%{!?go_compiler:golang}
BuildRequires:  btrfs-progs-devel
BuildRequires:  glib2-devel
BuildRequires:  glibc-static
BuildRequires:  go-md2man
BuildRequires:  gpgme-devel
BuildRequires:  libassuan-devel
BuildRequires:  pkgconfig(systemd)
BuildRequires:  device-mapper-devel
BuildRequires:  make
Requires(pre):  container-selinux
Requires:       skopeo-containers
Provides:       %{service_name} = %{version}-%{release}
Obsoletes:      ocid <= 0.3
Provides:       ocid = %{version}-%{release}
%if 0%{?rhel} >= 7
BuildRequires:  libseccomp-devel
Requires:   runc >= 1.0.0-10
%else
BuildRequires:  libseccomp-static
Recommends: containernetworking-plugins >= 0.5.2-2
Requires:   runc >= 2:1.0.0-10
%endif

%description
%{summary}

%prep
%setup -q -n %{name}
sed -i 's/\/usr\/local\/bin\/crio/\/usr\/bin\/crio/g' contrib/systemd/%{service_name}.service

%build
mkdir _build
pushd _build
mkdir -p src/%{provider}.%{provider_tld}/{%{project},opencontainers}
ln -s $(dirs +1 -l) src/%{import_path}
popd

ln -s vendor src
export GOPATH=$(pwd)/_build:$(pwd):$(pwd):%{gopath}
BUILDTAGS='selinux seccomp $(shell hack/btrfs_tag.sh) $(shell hack/libdm_tag.sh) containers_image_ostree_stub' make binaries docs

./bin/%{service_name} \
  --selinux=true \
  --storage-driver=overlay \
  --conmon /usr/libexec/%{service_name}/conmon \
  --cgroup-manager=systemd config > %{service_name}.conf

%install
make DESTDIR=%{buildroot} PREFIX=%{buildroot}%{_prefix} install.config install.systemd

# install binaries
install -dp %{buildroot}{%{_bindir},%{_libexecdir}/%{service_name}}
install -p -m 755 bin/%{service_name} %{buildroot}%{_bindir}
install -p -m 755 bin/crioctl %{buildroot}%{_bindir}
install -p -m 755 bin/conmon %{buildroot}%{_libexecdir}/%{service_name}
install -p -m 755 bin/pause %{buildroot}%{_libexecdir}/%{service_name}

install -dp %{buildroot}%{_sysconfdir}/cni/net.d
install -p -m 644 contrib/cni/10-crio-bridge.conf %{buildroot}%{_sysconfdir}/cni/net.d/100-crio-bridge.conf
install -p -m 644 contrib/cni/99-loopback.conf %{buildroot}%{_sysconfdir}/cni/net.d/200-loopback.conf

install -dp %{buildroot}%{_mandir}/man{5,8}
install -m 644 docs/*.5 %{buildroot}%{_mandir}/man5
install -m 644 docs/*.8 %{buildroot}%{_mandir}/man8

mkdir -p %{buildroot}%{_sharedstatedir}/containers

%check

%post
%systemd_post %{service_name}

%preun
%systemd_preun %{service_name}

%postun
%systemd_postun_with_restart %{service_name}

#define license tag if not already defined
%{!?_licensedir:%global license %doc}

%files
%license LICENSE
%doc README.md
%{_bindir}/%{service_name}
%{_bindir}/%{service_name}ctl
%{_mandir}/man5/%{service_name}.conf.5*
%{_mandir}/man8/%{service_name}.8*
%dir %{_sysconfdir}/%{service_name}
%config(noreplace) %{_sysconfdir}/%{service_name}/%{service_name}.conf
%config(noreplace) %{_sysconfdir}/%{service_name}/seccomp.json
%config(noreplace) %{_sysconfdir}/cni/net.d/100-%{service_name}-bridge.conf
%config(noreplace) %{_sysconfdir}/cni/net.d/200-loopback.conf
%config(noreplace) %{_sysconfdir}/crictl.yaml
%dir %{_libexecdir}/%{service_name}
%{_libexecdir}/%{service_name}/conmon
%{_libexecdir}/%{service_name}/pause
%{_unitdir}/%{service_name}.service
%{_unitdir}/cri-o.service
%{_unitdir}/%{service_name}-shutdown.service
%dir %{_sharedstatedir}/containers
%dir %{_datadir}/oci-umount
%dir %{_datadir}/oci-umount/oci-umount.d
%{_datadir}/oci-umount/oci-umount.d/%{service_name}-umount.conf

%changelog
* Thu Nov 02 2017 Lokesh Mandvekar <lsm5@fedoraproject.org> - 1.8.0-2.dev.testonly
- modify spec for running auto rpmbuild tests on PRs
- don't care about fedora changelog since this package is for PR testing only
