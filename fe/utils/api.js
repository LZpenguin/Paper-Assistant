const baseURL = 'http://doc_aid.bingyan.site'
function request(method,url,data,jwt) {
  return new Promise(function(resolve, reject) {
    let header = {
        'content-type': 'application/json',
    };
    if(jwt) {
      let token = wx.getStorageSync('token')
      header.Authorization = 'Bearer '+ token
    }
    wx.request({
        url: baseURL + url,
        method: method,
        data: method === 'POST' ? JSON.stringify(data) : data,
        header: header,
        success(res) {
            //请求成功
            if (res.data.status == 200) {
                resolve(res);
            } else {
                //其他错误
                console.log(res.errMsg);
            }
        },
        fail(err) {
            //请求失败
            reject(err)
        }
    })
  })
}

const API = {
  login: (data) => request('post','/user/login',data),
  document: () => request('GET','/document',{},true),
  detail: (id) => request('GET','/document/'+id,{},true),
  addFavFolder: (name) => request('POST','/folder/'+name,{},true),
  delFavFolder: (name) => request('DELETE','/folder/'+name,{},true),
  sortFavFolder: (names) => request('PUT','/folder',names,true),
  addFav: (name,ids) => request('POST','/fav/'+name,ids,true),
  delFav: (name,ids) => request('DELETE','/fav/'+name,ids,true),
  getFav: () => request('GET','/fav',{},true),
  sub: (name) => request('POST','/sub/'+name,{},true),
  unsub: (names) => request('DELETE','/sub',{names},true),
  getSubed: () => request('GET','/sub',{},true),
  mgzDetail: (name) => request('GET','/magazine/'+name,{},true),
  addKeyword: (keys) => request('POST','/keyword',keys,true),
  delKeyword: (keys) => request('DELETE','/keyword',keys,true),
  getKeyword: () => request('GET','/keyword',{},true)
}

module.exports = {
  API
}