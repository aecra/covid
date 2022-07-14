import Axios, { AxiosInstance } from 'axios';
import { RouteLocationNormalized } from 'vue-router';
import { buildQueryString, randomString, to } from './utils';

let Conf = {
  origin: '',
  authorization_endpoint: '',
  access_token_endpoint: '',
  refresh_token_endpoint: '',
  client_id: '',
  client_secret: '',
  redirect_uri: '',
  scope: '',
  state: '',
  response_type: '',
  api_base_url: '',
};

let RequestInstance: AxiosInstance;

if (localStorage.getItem('oauth_state') === null) {
  localStorage.setItem('oauth_state', randomString(16));
}
Conf.redirect_uri = window.location.origin + '/#/OAuth';

const Routes = [{ path: '/OAuth', name: 'OAuth', component: {} }];

const saveToken = (data: { access_token: string; refresh_token: string; expires_in: string; token_type: string }) => {
  localStorage.setItem('access_token', data.access_token);
  localStorage.setItem('refresh_token', data.refresh_token);
  localStorage.setItem('expires', String(Math.floor(Date.now() / 1000) + data.expires_in));
  localStorage.setItem('token_type', data.token_type);
};

const clearToken = () => {
  localStorage.removeItem('access_token');
  localStorage.removeItem('refresh_token');
  localStorage.removeItem('expires');
  localStorage.removeItem('token_type');
};

const verifyToken = async () => {
  if (!localStorage.getItem('access_token')) {
    // no access token
    return false;
  }
  if (Math.floor(Date.now() / 1000) > Number(localStorage.getItem('expires'))) {
    // access token expired
    // refresh token
    const [err] = await to(refreshAccessToken());
    return typeof err === null;
  }
  return true;
};

const getAccessToken = async (code: string) => {
  console.log('access_token request start');
  const res = await Axios({
    method: 'post',
    url: Conf.origin + Conf.access_token_endpoint,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      Authorization: 'Basic ' + Conf.client_id + ':' + Conf.client_secret,
    },
    data: buildQueryString({
      grant_type: 'authorization_code',
      code: code,
      redirect_uri: Conf.redirect_uri,
      client_id: Conf.client_id,
      client_secret: Conf.client_secret,
    }),
  });
  console.log('access_token request over');
  if (res.data.error || res.data.access_token.includes('error')) {
    console.log('access_token request is warning');
    return Promise.reject('get access token error');
  }
  console.log('access_token request is ok');
  saveToken(res.data);
};

const refreshAccessToken = () => {
  return new Promise((resolve, reject) => {
    if (localStorage.getItem('refresh_token')) {
      Axios({
        method: 'post',
        url: Conf.origin + Conf.refresh_token_endpoint,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
          Authorization: 'Basic ' + btoa(Conf.client_id + ':' + Conf.client_secret),
        },
        data: buildQueryString({
          grant_type: 'refresh_token',
          refresh_token: localStorage.getItem('refresh_token') as string,
          client_id: Conf.client_id,
        }),
      })
        .then((res) => {
          if (res.data.error || res.data.access_token.includes('error')) {
            reject(Error('refresh access token error'));
            return;
          }
          saveToken(res.data);
          resolve('refresh success');
        })
        .catch(() => {
          reject(Error('refresh access token error'));
        });
    } else {
      reject(Error('refresh_token not found'));
    }
  });
};

const redirectToLogin = () => {
  clearToken();
  window.location.replace(
    Conf.origin +
      Conf.authorization_endpoint +
      '?' +
      buildQueryString({
        client_id: Conf.client_id,
        response_type: Conf.response_type,
        scope: Conf.scope,
        state: localStorage.getItem('oauth_state') as string,
        redirect_uri: Conf.redirect_uri,
      })
  );
};

const OAuth = (data: Record<string, unknown>) => {
  if (data.state != localStorage.getItem('oauth_state') || data.error || !data.code) {
    redirectToLogin();
  }
  getAccessToken(<string>data.code)
    .then(() => window.location.replace('/'))
    .catch(() => {
      console.log('get access token error');
      redirectToLogin();
    });
};

const NavigationGuard = async (to: RouteLocationNormalized) => {
  if (to.path !== '/OAuth' && !(await verifyToken())) {
    return '/OAuth';
  }
  if (to.path === '/OAuth') {
    OAuth(to.query as { state: string; code: string; error: string });
    return false;
  }
};

const Init = (conf: Record<string, unknown>) => {
  Conf = { ...Conf, ...conf };
};

const Request = () => {
  if (!RequestInstance) {
    RequestInstance = Axios.create({
      baseURL: Conf.api_base_url,
      timeout: 3000,
      responseType: 'json',
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('access_token'),
      },
    });
    RequestInstance.interceptors.request.use((config) => {
      if (!verifyToken()) {
        redirectToLogin();
      }
      return config;
    });
    // if state is 4xx, redirect to OAuth
    RequestInstance.interceptors.response.use((res) => {
      if (res.status >= 400 && res.status < 500) {
        redirectToLogin();
      }
      return res;
    });
  }
  return RequestInstance;
};

export default {
  NavigationGuard,
  Init,
  Request,
  Routes,
};
