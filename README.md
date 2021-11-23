# Euler robot function

| 功能                                          | 差异描述                                                     | 状态   |
| --------------------------------------------- | ------------------------------------------------------------ | ------ |
| welcome                                       | PR创建时，weclome消息需展示所属sig组，并展示maintainer。e.g：<br />Hi ***ssttkx***, welcome to the openEuler Community.<br/>I'm the Bot here serving you. You can find the instructions on how to interact with me at<br/>https://gitee.com/openeuler/community/blob/master/en/sig-infrastructure/command.md.<br/>If you have any questions, please contact the SIG: [sig-QA](https://gitee.com/openeuler/community/tree/master/sig/sig-QA), and any of the maintainers: ***[@lemon.higgins ](https://gitee.com/open_euler/dashboard/members/lemon-higgins)***, ***[@ltx ](https://gitee.com/open_euler/dashboard/members/lutianxiong)***, ***[@Fengguang ](https://gitee.com/open_euler/dashboard/members/wu_fengguang)***, ***[@Charlie_Li ](https://gitee.com/open_euler/dashboard/members/Charlie_li)***, ***[@wubodong ](https://gitee.com/open_euler/dashboard/members/walkingwalk)***, ***[@KuhnChen ](https://gitee.com/open_euler/dashboard/members/kuhnchen18)***, ***[@xxxxxxxxxxxx ](https://gitee.com/speacher)***, ***[@将进酒杯莫停 ](https://gitee.com/open_euler/dashboard/members/rigorous)***. | 需开发 |
| /retest                                       | pr 有新的commit提交，会自动添加/retest评论                   | 需开发 |
| /review-tool                                  | 单世史的python外部插件，prow hook-agent插件调用该插件        | 需讨论 |
| 自动添加sig/*标签                             | issue或PR创建时会根据仓库所属的sig组自动添加对应的sig标签    | 需开发 |
| 检查PR是否设置Reviewer                        | PR创建时，或检测PR创建者是否设置reviewer，未创建给出提示     | 需开发 |
| lgtm/approve/merge                            | merge支持设置分支冻结，冻结的分支PR不能自动合入，只有冻结配置特定用户有权限合入<br />/approve [cancel]指令的使用权限不再针对类似community这种特殊仓下sigs组中的maintainer做判断 | 需开发 |
| watch special file change and do some options | [code-line](https://gitee.com/openeuler/ci-bot/blob/master/pkg/cibot/pullrequest.go#L194) | 待讨论 |

