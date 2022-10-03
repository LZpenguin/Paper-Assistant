// components/search/search.js
Component({
  properties: {
    title: {
      type: String,
      default: '搜索'
    },
    enable: {
      type: Boolean,
      default: true
    },
    jump: {
      type: String
    }
  },
  data: {
    search: false
  },
  methods: {
    toggle: function(e) {
      this.setData({
        search: e.currentTarget.dataset.search==0?false:true
      })
    },
    bindInput(e) {
      this.setData({
        input: e.detail.value
      })
    },
    clear(e) {
      this.setData({
        input: ''
      })
    },
    search(e) {
      this.triggerEvent('searchValue',{value: this.data.input.trim()})
    },
    jump(e) {
      wx.navigateTo({
        url: e.currentTarget.dataset.jump,
      })
    }
  }
})
