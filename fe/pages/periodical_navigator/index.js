// pages/periodical_navigator/index.js
Page({
  data: {
    naviList: [],
    checked: 0
  },
  showClass(e) {
    this.setData({
      checked: e.currentTarget.dataset.index
    })
  },
  jump_periodical_class(e) {
    wx.navigateTo({
      url: '/pages/periodical_class/index?navi='+e.currentTarget.dataset.name+'&class='+e.currentTarget.dataset.class,
    })
  },
  getData() {
    var that = this;
    wx.getStorage({
      key: 'token',
      success: res => {
        this.setData({
          'token': res.data
        })
        wx.request({
          url: 'https://api.hust.online/doc-aid-test/magazine',
          success: function(rel) {
            that.setData({
              naviList: rel.data.data.slice(1,)
            })
          },
          failed: function (err) {
            setTimeout(() => {
              this.getData()
            }, 200);
          },
          header: {
            'Authorization': 'Bearer '+ res.data
          }
        })
      }
    })
  },
  onLoad() {
    wx.getStorage({
      key: 'naviList',
      success: res => {
        this.setData({
          naviList: res.data.slice(1,)
        })
      }
    })
    this.getData()
  }
})