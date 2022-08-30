import Axios, {AxiosInstance, AxiosRequestConfig} from 'axios';
import {RouteLocationNormalized} from "vue-router";

let RequestInstance: AxiosInstance;

const randomString = (e: number): string => {
  e = e || 32;
  const t = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890';
  const a = t.length;
  let n = '';
  for (let i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
  return n;
};

const GetRequest = () => {
  if (!RequestInstance) {
    RequestInstance = Axios.create({
      responseType: 'json',
      headers: {
        'Authorization': 'Bearer ' + localStorage.getItem('token'),
        'Cache-Control': 'no-cache',
      },
    });
    RequestInstance.interceptors.request.use((config: AxiosRequestConfig) => {
      if (/get/i.test(config?.method as string)) { //判断get请求
        config.params  =  config.params || {};
        config.params.t = randomString(16);
      }
      return config;
    }, error => {
      return Promise.reject(error);
    })
  }
  return RequestInstance;
};

interface RegisterRequest {
  username: string,
  email: string,
  password: string,
  confirmation_password: string,
}

const register = async (registerRequest: RegisterRequest) => {
  return Axios.post('/api/auth/register', registerRequest)
}

interface LoginRequest {
  username: string,
  password: string,
}

const login = (loginRequest: LoginRequest) => {
  return Axios.post('/api/auth/login', loginRequest)
}

const online = async () => {
  return await GetRequest().post('/api/auth/online').then((res) => {
    return res.status === 200;
  }).catch(() => {
    return false;
  });
}

const guard = async (to: RouteLocationNormalized) => {
  if ((to.path !== '/login' && to.path !== '/register') && !(await online())) {
    localStorage.removeItem("token")
    return '/login';
  }
};

export default {
  GetRequest,
  register,
  login,
  guard,
};
