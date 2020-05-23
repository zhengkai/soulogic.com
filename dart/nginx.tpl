server {

	server_name %DOMAIN%;

	listen [::]:443 ssl http2;

	ssl_certificate     ssl.d/%HOSTNAME%.9farm.com.crt;
    ssl_certificate_key ssl.d/%HOSTNAME%.9farm.com.key;

	add_header Strict-Transport-Security "max-age=99999999; includeSubDomains; preload";

	add_header X-Frame-Options SAMEORIGIN;
	add_header X-Content-Type-Options nosniff;
	add_header X-XSS-Protection "1; mode=block";

	access_log %LOG%/access.log;
	error_log  %LOG%/error.log;

	root /www/%NAME%/client/dist/%NAME%;

	error_page 404 /index.html;

	location /webhook {
		add_header "Access-Control-Allow-Origin" "*";
		proxy_pass http://127.0.0.1:9025;
	}

	location / {
		try_files $uri $uri/ /index.html;
	}
}

server {

	server_name %DOMAIN%

	listen [::]:80;

	location / {
		return 301 https://$host$request_uri;
	}
}
