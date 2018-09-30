```
1) 错误点 php-pfm的执行 /var/www/html 要和nginx的容器的目录吗要相同不然，无法找到项目的目录

server {
    listen       80;
    server_name  localhost;

    #charset koi8-r;
    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /data/web;
        index  index.html index.htm index.php;
		 try_files $uri $uri/ =404;
    }
	autoindex on;
	autoindex_exact_size off;
	autoindex_localtime on;

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /data/web;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    location ~ \.php$ {
        root           /data/web;
        fastcgi_pass   172.17.0.4:9000; # PHP-fpm的容器IP docker inspect php 查看
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }

    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    location ~ /\.ht {
        deny  all;
    }
}



```
