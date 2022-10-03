const { API } = require("../../utils/api");

// pages/my_subscribe/index.js
Page({
  data: {
    periodicalList: []
  },
  getData() {
    var that = this;
    API.getSubed().then(res => {
      that.setData({
        periodicalList: res.data.data
      })
    })
    API.getKeyword().then(res => {
      that.setData({
        keywordList: res.data.data
      })
    })
  },
  onShow(e) {
    this.getData()
  },
  jump_periodical_navi(e) {
    wx.navigateTo({
      url: '/pages/periodical_navigator/index',
    })
  },
  jump_add_keywords(e) {
    wx.navigateTo({
      url: '/pages/add_keywords/index',
    })
  }
})