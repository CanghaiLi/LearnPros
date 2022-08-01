# CICD: Git + Jenkins + Harbor + Docker

### 1. 部署Git服务

git-server的设置操作：

```bash
yum install -y git git-core gitweb
mkdir -p /git-root
cd /git-root && git init --bare shell.git
# --bare 创建git裸库，适用于origin远程git服务，不允许用户进行任何git操作
chown -R root:root shell.git # 权限更改

# 秘钥设置
cd ~ && ssh-keygen -t rsa
cd .ssh && cp id_rsa.pub authorized_keys

```

