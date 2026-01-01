import os

# 设置中文
LANGUAGE_CODE = 'zh-hans'

# 设置数据库
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'vul_server',
        'USER': 'root',
        'PASSWORD': 'root',
        'HOST': os.getenv("DB_HOST", "db"),
        'PORT': 3306,
    }
}