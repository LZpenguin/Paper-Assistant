let {API} = require('./utils/api')
// app.js
App({
  onLaunch() {
    wx.login({
      timeout: 1000, 
      success: res => {
        API.login({
          code: res.code
        }).then(res => {
          wx.setStorageSync('token', res.data.data.token)
          API.getFav().then(res => {
            if(res.data.data.length == 0 ) {
              API.addFavFolder('默认收藏')
            }
          })
        },err => {
          console.log(err);
        })
      }
    })
  }
})
