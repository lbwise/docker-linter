WORKDIR /usr/local/app

# Install the application dependencies
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

FROM python:3.12

# Copy in the source code
EXPOSE '7000-8000'
EXPOSE

# Setup an app user so the container doesn't run as the root user
RUN ''
USER ''

CMD ["service", "--config", "/etc/service.conf"]

# order of commands dont matter - thus per command linting
# number of arguments for some commands