// components/cover&name&if/cover.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    'if': {
      type: Number,
      value: '0.0'
    },
    'name': {
      type: String,
      value: '期刊名称'
    },
    'img': {
      type: String,
      value: '../../icons/cover_rect.png'
    }
  },
  data: {

  },
  methods: {
    jump_periodical_detail(e) {
      wx.navigateTo({
        url: '/pages/periodical_detail/index?name='+e.currentTarget.dataset.magazine,
      })
    }
  }
})
