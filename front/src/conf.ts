const conf = {
  OAuth: {
      origin: 'https://oauth.aecra.cn',
      authorization_endpoint: '/login/oauth/authorize',
      access_token_endpoint: '/api/login/oauth/access_token',
      refresh_token_endpoint: '/api/login/oauth/refresh_token',
      client_id: '4278f4c62cb5764f644e',
      client_secret: 'e7170c761c47483e6f57a0e94c94248f280a573f',
      scope: 'openid email',
      response_type: 'code',
      api_base_url: 'https://api3.aecra.cn/covid/',
  },
}

export default conf;