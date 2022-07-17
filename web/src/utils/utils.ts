const buildQueryString = (params: Record<string, string>) => {
  let queryString = '';
  for (const key in params) {
    if (Object.prototype.hasOwnProperty.call(params, key)) {
      queryString += encodeURIComponent(key) + '=' + encodeURIComponent(params[key]) + '&';
    }
  }
  return queryString;
};

const randomString = (e: number) => {
  e = e || 32;
  const t = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890';
  const a = t.length;
  let n = '';
  for (let i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
  return n;
};

const to = (promise: Promise<unknown>) => {
  return promise
    .then((data) => {
      return [null, data];
    })
    .catch((err) => [err]);
};

export { buildQueryString, randomString, to };
