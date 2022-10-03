// pages/user/index.js
Page({
  data: {

  },
  jump_my_subscribe(e) {
    wx.navigateTo({
      url: '/pages/my_subscribe/index',
    })
  },
  jump_shared(e) {
    wx.navigateTo({
      url: '/pages/shared/index',
    })
  },
  jump_aboutUs(e) {
    wx.navigateTo({
      url: '/pages/aboutUs/index',
    })
  },
  onShow() {
    wx.getStorage({
      key: 'userInfo',
      success: res => {
        this.setData({
          userInfo: res.data
        })
      }
    })
    wx.getStorage({
      key: 'fromIndex',
      success: res => {
        if(res.data) {
          wx.setStorage({
            key: 'fromIndex',
            data: false
          })
          wx.navigateTo({
            url: '/pages/my_subscribe/index',
          })
        }
      }
    })
  }
})