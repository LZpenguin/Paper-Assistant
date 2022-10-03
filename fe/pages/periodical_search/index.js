// pages/periodical_seach/index.js
Page({
  data: {
    periodicalList: [],
    result: []
  },
  getData(name) {
    var that = this;
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/magazine/'+name,
          success: (rel) => {
            this.setData({
              result: [rel.data.data]
            })
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  },
  getValue(e) {
    let reg = new RegExp(e.detail.value);
    this.getData(e.detail.value)
  }
})