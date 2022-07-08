import Axios from 'axios';
import { RouteLocationNormalized } from 'vue-router';
import { buildQueryString } from './utils';

let config = {
  authorization_endpoint: 'https://oauth.aecra.cn/login/oauth/authorize',
  token_endpoint: 'https://oauth.aecra.cn/api/login/oauth/access_token',
  refresh_token_endpoint: 'https://oauth.aecra.cn/api/login/oauth/refresh_token',
  client_id: '4278f4c62cb5764f644e',
  client_secret: 'e7170c761c47483e6f57a0e94c94248f280a573f',
  redirect_uri: window.location.protocol + '//' + window.location.host + '/#/Oauth',
  scope: 'openid email profile address phone offline_access',
  state: 'viursdhgnbiseyagb',
  response_type: 'code',
  login_url: 'https://oauth.aecra.cn/login/oauth/authorize',
};

//TODO: access_token and refresh_token expired
let verifyToken = async () => {
  if (!localStorage.getItem('access_token')) {
    // no access token
    return false;
  }
  if (Math.floor(Date.now() / 1000) > Number(localStorage.getItem('expires'))) {
    // access token expired
    // refresh token
    const res = await refrershAccessToken();
    if (res === 'refresh success') {
      return true;
    }
    return false;
  }
  return true;
};

let getAccessToken = async (code: string) => {
  console.log('access_token request start');
  const res = await Axios({
    method: 'post',
    url: config.token_endpoint,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      Authorization: 'Basic ' + config.client_id + ':' + config.client_secret,
    },
    data: buildQueryString({
      grant_type: 'authorization_code',
      code: code,
      redirect_uri: config.redirect_uri,
      client_id: config.client_id,
      client_secret: config.client_secret,
    }),
  });
  console.log('access_token request over');
  if (res.data.error || res.data.access_token.includes('error')) {
    console.log('access_token request is worning');
    return Promise.reject('get access token error');
  }
  console.log('access_token request is ok');
  localStorage.setItem('access_token', res.data.access_token);
  localStorage.setItem('refresh_token', res.data.refresh_token);
  localStorage.setItem('expires', String(Math.floor(Date.now() / 1000) + res.data.expires_in));
  localStorage.setItem('token_type', res.data.token_type);
};

let refrershAccessToken = () => {
  return new Promise((resolve, reject) => {
    if (localStorage.getItem('refresh_token')) {
      Axios({
        method: 'post',
        url: config.refresh_token_endpoint,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
          Authorization: 'Basic ' + btoa(config.client_id + ':' + config.client_secret),
        },
        data: buildQueryString({
          grant_type: 'refresh_token',
          refresh_token: localStorage.getItem('refresh_token') as string,
          client_id: config.client_id,
        }),
      })
        .then((res) => {
          if (res.data.error || res.data.access_token.includes('error')) {
            reject('refresh token error');
            return;
          }
          localStorage.setItem('access_token', res.data.access_token);
          localStorage.setItem('refresh_token', res.data.refresh_token);
          localStorage.setItem('expires', String(Math.floor(Date.now() / 1000) + res.data.expires_in));
          localStorage.setItem('token_type', res.data.token_type);
          resolve('refresh success');
        })
        .catch((err) => {
          reject('refresh_token request error');
        });
    } else {
      reject('no refresh_token');
    }
  });
};

let redirectToLogin = () => {
  window.location.replace(
    config.login_url +
      '?' +
      buildQueryString({
        client_id: config.client_id,
        response_type: config.response_type,
        scope: config.scope,
        state: config.state,
        redirect_uri: config.redirect_uri,
      })
  );
};

let Oauth = (data: any) => {
  if (data.state != config.state) {
    redirectToLogin();
  } else if (data.error) {
    redirectToLogin();
  } else {
    if (data.code) {
      getAccessToken(data.code)
        .then(() => {
          window.location.href = '/';
        })
        .catch(() => {
          redirectToLogin();
        });
    } else {
      redirectToLogin();
    }
  }
};

let navigationGuard = async (to: RouteLocationNormalized) => {
  if (to.path !== '/Oauth' && !(await OauthService.verifyToken())) {
    return '/Oauth';
  }
  if (to.path === '/Oauth') {
    OauthService.Oauth(to.query);
    return false;
  }
};

const OauthService = {
  Oauth,
  verifyToken,
  navigationGuard,
};

export default OauthService;
