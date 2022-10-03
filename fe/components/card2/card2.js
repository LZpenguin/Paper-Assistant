// components/card2/card2.js
Component({
  properties: {
    name: {
      type: String,
      value: '期刊名称'
    },
    enName: {
      type: String,
      value: 'English Name'
    },
    sub: {
      type: Boolean,
      value: false
    },
    cycle: {
      type: String
    },
    fif: {
      type: Number
    },
    zif: {
      type: Number
    },
    img: {
      type: String
    }
  },
  methods: {
    add_sub: function(e) {
      wx.getStorage({
        key: 'token',
        success: res => {
          wx.request({
            url: 'https://api.hust.online/doc-aid-test/sub/'+e.currentTarget.dataset.name,
            method: 'post',
            header: {
              'Authorization': 'Bearer '+ res.data
            },
            success: () => {
              this.setData({
                sub: true
              })
            }
          })
        }
      })
    },
    delete_sub: function(e) {
      wx.getStorage({
        key: 'token',
        success: res => {
          wx.request({
            url: 'https://api.hust.online/doc-aid-test/sub',
            method: 'delete',
            data: [e.currentTarget.dataset.name],
            header: {
              'Authorization': 'Bearer '+ res.data
            },
            success: () => {
              this.setData({
                sub: false
              })
            }
          })
        }
      })
    },
    jump_periodical_detail(e) {
      wx.navigateTo({
        url: '/pages/periodical_detail/index?name='+e.currentTarget.dataset.magazine,
      })
    }
  }
})
