# Copyright 2015 The Kubernetes Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


FROM reg.docker.alibaba-inc.com/alibase/alios7u2:latest as builder

COPY . /

RUN /build.sh

# Use a multi-stage build
FROM reg.docker.alibaba-inc.com/alibase/alios7u2:latest

ENV PATH=$PATH:/usr/local/luajit/bin:/usr/local/tengine/sbin:/usr/local/tengine/bin

ENV LUA_PATH="/usr/local/share/luajit-2.1.0-beta3/?.lua;/usr/local/share/lua/5.1/?.lua;/usr/local/lib/lua/?.lua;;"
ENV LUA_CPATH="/usr/local/lib/lua/?/?.so;/usr/local/lib/lua/?.so;;"

COPY --from=builder /usr/local /usr/local
COPY --from=builder /opt /opt
COPY --from=builder /etc/nginx /etc/nginx

RUN rm -rf /usr/local/nginx/sbin/nginx && chmod 6755 /usr/local/tengine/sbin/tengine && chown root:admin /usr/local/tengine/sbin/tengine

RUN rpm -rebuilddb && yum install -y GeoIP GeoIP-devel

RUN rpm -rebuilddb && yum install -y -b current t-dumb-init && ln -s /usr/local/bin/dumb-init /usr/bin/dumb-init

RUN test ! -d /home/coredump && mkdir /home/coredump && chmod 777 /home/coredump

EXPOSE 80 443

CMD ["tengine", "-g", "daemon off;"]
