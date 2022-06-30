import Axios from 'axios';
import OauthService from './oauth';

const instance = Axios.create({
  baseURL: 'https://api3.aecra.cn/covid/',
  timeout: 1000,
  responseType: 'json',
  headers: { 'Content-Type': 'application/json;charset=utf-8' },
});

instance.interceptors.request.use(
  function (config) {
    // 在发送请求之前做些什么
    if (!OauthService.verifyToken()) {
      throw new Axios.Cancel('Operation canceled by the user.');
    }
    (config as any).headers['Authorization'] = 'Bearer ' + localStorage.getItem('access_token');
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

instance.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    if(error.response.status === 401){
      console.log('401');
    }
    return Promise.reject(error);
  }
);

export default instance;
