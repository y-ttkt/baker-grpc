FROM ruby:3.0.5

RUN apt-get update -qq \
 && curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
 && apt-get install -y nodejs \
 && npm install --global yarn

WORKDIR /app

COPY Gemfile Gemfile.lock ./

RUN bundle install

COPY . .

CMD ["bash", "-lc", "rails server -b 0.0.0.0 -p 3000"]
