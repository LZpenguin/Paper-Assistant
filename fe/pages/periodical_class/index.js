// pages/periodical_class/index.js
Page({
  data: {
    periodicalList: []
  },
  onLoad(e) {
    this.getData(e.navi,e.class)
    wx.setNavigationBarTitle({
      title: e.class
    })
  },
  getData(issue,topic) {
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/magazine/'+issue+'/'+topic,
          success: rel => {
            this.setData({
              periodicalList: rel.data.data
            })
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  }
})