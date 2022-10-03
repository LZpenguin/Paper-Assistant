const { API } = require("../../utils/api");

// pages/periodical_detail/index.js
Page({
  data: {
    periodical: {},
  },
  getData(name) {
    var that = this;
    API.mgzDetail(name).then(res => {
      that.setData({
        periodical: res.data.data
      })
    })
  },
  add_sub: function(e) {
    let name = e.currentTarget.dataset.name
    let that = this
    API.sub(name).then(res => {
      let periodical = that.data.periodical;
        periodical.sub = true;
        that.setData({
          periodical
        })
    })
  },
  delete_sub: function(e) {
    let names = [e.currentTarget.dataset.name]
    let that = this
    API.unsub(names).then(res => {
      let periodical = that.data.periodical;
      periodical.sub = false;
      that.setData({
        periodical
      })
    })
  },
  onLoad(e) {
    this.getData(e.name)
  }
})