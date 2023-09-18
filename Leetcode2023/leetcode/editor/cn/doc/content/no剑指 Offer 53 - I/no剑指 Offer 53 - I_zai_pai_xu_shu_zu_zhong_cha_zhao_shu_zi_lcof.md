<p>统计一个数字在排序数组中出现的次数。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> nums = [<span><code>5,7,7,8,8,10]</code></span>, target = 8
<strong>输出:</strong> 2</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre>
<strong>输入:</strong> nums = [<span><code>5,7,7,8,8,10]</code></span>, target = 6
<strong>输出:</strong> 0</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>0 &lt;= nums.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>-10<sup>9</sup>&nbsp;&lt;= nums[i]&nbsp;&lt;= 10<sup>9</sup></code></li> 
 <li><code>nums</code>&nbsp;是一个非递减数组</li> 
 <li><code>-10<sup>9</sup>&nbsp;&lt;= target&nbsp;&lt;= 10<sup>9</sup></code></li> 
</ul>

<p>&nbsp;</p>

<p><strong>注意：</strong>本题与主站 34 题相同（仅返回值不同）：<a href="https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/">https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/</a></p>

<details><summary><strong>Related Topics</strong></summary>数组 | 二分查找</details><br>

<div>👍 446, 👎 0<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.gitee.io/article/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.github.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 和 [递归算法专题课](https://aep.xet.tech/s/3YGcq3) 限时附赠网站会员，[新版刷题打卡挑战](https://labuladong.gitee.io/algo/challenge/) 上线！**



<p><strong><a href="https://labuladong.github.io/article/slug.html?slug=zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题考察二分搜索查找 `target` 的左右边界，和 [34. 在排序数组中查找元素的第一个和最后一个位置](/problems/find-first-and-last-position-of-element-in-sorted-array) 有些类似，用二分搜索找到左右边界的索引，就可以判断重复出现的次数了。

**详细题解：[我写了首诗，把二分搜索算法变成了默写题](https://labuladong.github.io/article/fname.html?fname=二分查找详解)**

**标签：[二分搜索](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)，[数组](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

## 解法代码

提示：🟢 标记的是我写的解法代码，🤖 标记的是 chatGPT 翻译的多语言解法代码。如有错误，可以 [点这里](https://github.com/labuladong/fucking-algorithm/issues/1113) 反馈和修正。

<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// 注意：cpp 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left_index = left_bound(nums, target);
        if (left_index == -1) {
            return 0;
        }
        int right_index = right_bound(nums, target);
        // 根据左右边界即可推导出元素出现的次数
        return right_index - left_index + 1;
    }

    int left_bound(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        // 搜索区间为 [left, right]
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                // 搜索区间变为 [mid+1, right]
                left = mid + 1;
            } else if (nums[mid] > target) {
                // 搜索区间变为 [left, mid-1]
                right = mid - 1;
            } else if (nums[mid] == target) {
                // 收缩右侧边界
                right = mid - 1;
            }
        }
        // 检查出界情况
        if (left >= nums.size() || nums[left] != target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/2.jpg) */
            return -1;
        }
        return left;
    }

    int right_bound(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid + 1;
            } else if (nums[mid] > target) {
                right = mid - 1;
            } else if (nums[mid] == target) {
                // 这里改成收缩左侧边界即可
                left = mid + 1;
            }
        }
        // 这里改为检查 right 越界的情况，见下图
        if (right < 0 || nums[right] != target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/4.jpg) */
            return -1;
        }
        return right;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left_index = self.left_bound(nums, target)
        if left_index == -1:
            return 0
        right_index = self.right_bound(nums, target)
        # 根据左右边界即可推导出元素出现的次数
        return right_index - left_index + 1

    def left_bound(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        # 搜索区间为 [left, right]
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                # 搜索区间变为 [mid+1, right]
                left = mid + 1
            elif nums[mid] > target:
                # 搜索区间变为 [left, mid-1]
                right = mid - 1
            elif nums[mid] == target:
                # 收缩右侧边界
                right = mid - 1
        # 检查出界情况
        if left >= len(nums) or nums[left] != target: # <extend up -300>![](https://labuladong.github.io/pictures/二分查找/2.jpg) #
            return -1
        return left

    def right_bound(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                # 这里改成收缩左侧边界即可
                left = mid + 1
        # 这里改为检查 right 越界的情况，见下图
        if right < 0 or nums[right] != target: # <extend up -300>![](https://labuladong.github.io/pictures/二分查找/4.jpg) #
            return -1
        return right
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public int search(int[] nums, int target) {
        int left_index = left_bound(nums, target);
        if (left_index == -1) {
            return 0;
        }
        int right_index = right_bound(nums, target);
        // 根据左右边界即可推导出元素出现的次数
        return right_index - left_index + 1;
    }

    int left_bound(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        // 搜索区间为 [left, right]
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                // 搜索区间变为 [mid+1, right]
                left = mid + 1;
            } else if (nums[mid] > target) {
                // 搜索区间变为 [left, mid-1]
                right = mid - 1;
            } else if (nums[mid] == target) {
                // 收缩右侧边界
                right = mid - 1;
            }
        }
        // 检查出界情况
        if (left >= nums.length || nums[left] != target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/2.jpg) */
            return -1;
        }
        return left;
    }

    int right_bound(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] < target) {
                left = mid + 1;
            } else if (nums[mid] > target) {
                right = mid - 1;
            } else if (nums[mid] == target) {
                // 这里改成收缩左侧边界即可
                left = mid + 1;
            }
        }
        // 这里改为检查 right 越界的情况，见下图
        if (right < 0 || nums[right] != target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/4.jpg) */
            return -1;
        }
        return right;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

// 二分查找
// 在排序数组中查找元素的第一个和最后一个位置
func search(nums []int, target int) int {
    // 元素第一次出现的位置
    leftIndex := leftBound(nums, target)
    if leftIndex == -1 {
        return 0
    }
    // 元素最后一次出现的位置
    rightIndex := rightBound(nums, target)
    // 根据左右边界即可推导出元素出现的次数
    return rightIndex - leftIndex + 1
}

// 二分查找，查找元素第一次出现的位置
func leftBound(nums []int, target int) int {
    left, right := 0, len(nums)-1
    // 搜索区间为 [left, right]
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            // 搜索区间变为 [mid+1, right]
            left = mid + 1
        } else if nums[mid] > target {
            // 搜索区间变为 [left, mid-1]
            right = mid - 1
        } else if nums[mid] == target {
            // 收缩右侧边界
            right = mid - 1
        }
    }
    // 检查出界情况
    if left >= len(nums) || nums[left] != target {
        // 该元素在数组中不存在的情况
        return -1
    }
    return left
}

// 二分查找，查找元素最后一次出现的位置
func rightBound(nums []int, target int) int {
    left, right := 0, len(nums)-1
    // 搜索区间为 [left, right]
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            // 搜索区间变为 [mid+1, right]
            left = mid + 1
        } else if nums[mid] > target {
            // 搜索区间变为 [left, mid-1]
            right = mid - 1
        } else if nums[mid] == target {
            // 收缩左侧边界
            left = mid + 1
        }
    }
    // 检查出界情况
    if right < 0 || nums[right] != target {
        // 该元素在数组中不存在的情况
        return -1
    }
    return right
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
  function left_bound(nums, target) {
    let left = 0
    let right = nums.length - 1
    // 搜索区间为 [left, right]
    while (left <= right) {
      let mid = left + Math.floor((right - left) / 2)
      if (nums[mid] < target) {
        // 搜索区间变为 [mid+1, right]
        left = mid + 1
      } else if (nums[mid] > target) {
        // 搜索区间变为 [left, mid-1]
        right = mid - 1
      } else if (nums[mid] === target) {
        // 收缩右侧边界
        right = mid - 1
      }
    }
    // 检查出界情况
    if (left >= nums.length || nums[left] !== target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/2.jpg) */
      return -1;
    }
    return left;
  }

  function right_bound(nums, target) {
    let left = 0
    let right = nums.length - 1
    while (left <= right) {
      let mid = left + Math.floor((right - left) / 2)
      if (nums[mid] < target) {
        left = mid + 1
      } else if (nums[mid] > target) {
        right = mid - 1
      } else if (nums[mid] === target) {
        // 这里改成收缩左侧边界即可
        left = mid + 1
      }
    }
    // 这里改为检查 right 越界的情况，见下图
    if (right < 0 || nums[right] !== target) {/**<extend up -300>![](https://labuladong.github.io/pictures/二分查找/4.jpg) */
      return -1;
    }
    return right;
  }

  let left_index = left_bound(nums, target);
  if (left_index === -1) {
    return 0;
  }
  let right_index = right_bound(nums, target);
  // 根据左右边界即可推导出元素出现的次数
  return right_index - left_index + 1;
};
```

</div></div>
</div></div>

**类似题目**：
  - [34. 在排序数组中查找元素的第一个和最后一个位置](/problems/find-first-and-last-position-of-element-in-sorted-array)
  - [704. 二分查找](/problems/binary-search)

</details>
</div>

