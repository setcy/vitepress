package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve() {
	ginServer := gin.New()

	ginServer.Use(cors())

	serveContent(ginServer)

	err := ginServer.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func serveContent(ginServer *gin.Engine) {
	ginServer.GET("/_meta", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": []gin.H{
				{
					"text": "Home",
					"items": []gin.H{
						{
							"text": "Child",
							"link": "/child",
						},
					},
				},
				{
					"text": "About",
					"link": "/about",
				},
				{
					"text": "Contact",
					"link": "/contact",
				},
			},
		})
	})
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"content": `
<h1 class="toc-header" id="医疗服务"><a href="#医疗服务" class="toc-anchor">¶</a> <strong>医疗服务</strong></h1>
<p>入学后，学校会统一安排，进行一次体检（核心项目为血常规和胸扫），并且以班级为单位缴纳保险（关于医疗保险的相关信息，参见 转诊和医疗保险政策 ）的费用。</p>
<p>学校内的校医院提供最基本的医疗服务，包含全科门诊、基础检查、基本药品。</p>
<p>学校附近包含两个卫生服务中心（下沙街道卫生服务中心、白杨街道卫生服务中心）提供基本医疗服务（如疫苗注射等）。</p>
<p>学校附近包含两个三甲医院（浙江省中医院（浙江中医药大学附属第一医院）钱塘院区、浙江大学医学院附属邵逸夫医院钱塘院区）（简称中医院、邵逸夫），满足绝大部分医疗需求，其中中医院相对近。</p>
<h2 class="toc-header" id="校医院"><a href="#校医院" class="toc-anchor">¶</a> 校医院</h2>
<h3 class="toc-header" id="开放时间"><a href="#开放时间" class="toc-anchor">¶</a> 开放时间</h3>
<p>工作日 8:00—20:00</p>
<p>节假日 8:00–11:30 13:30—16:00</p>
<h3 class="toc-header" id="就诊"><a href="#就诊" class="toc-anchor">¶</a> 就诊</h3>
<p><strong>携带材料</strong></p>
<ul>
<li>一卡通</li>
<li>医保卡 或 电子医保凭证 或 市民卡</li>
</ul>
<p><strong>就诊流程</strong></p>
<ul>
<li>挂号</li>
<li>看病</li>
<li>付费&amp;取药（医保会直接报销对应部分）</li>
</ul>
<p><strong>就诊流程</strong></p>
<p><strong>一、自费挂号</strong><br>
凭一卡通→挂号→看病→付费→取号</p>
<p><strong>二、医保挂号</strong></p>
<p><strong>①办理过医保卡/市民卡</strong><br>
凭医保卡/市民卡/电子医保凭证、一卡通→挂号→看病→付费→取药</p>
<p><strong>②未办理过医保卡/市民卡</strong><br>
凭一卡通、医保证历本→挂号→看病→付费→取药</p>
<p>说明：根据杭州市医疗保障局说明，2020年4月1日起，全面推广应用电子病历，医疗保障部门不再发放纸质证历本。因此多数情况下为①情况，考虑到可能涉及到跨市跨地区等复杂情况，保留②的说明，<strong>有效性请咨询校医院和属地医保机构</strong>。</p>
<h2 class="toc-header" id="转诊和医疗保险政策"><a href="#转诊和医疗保险政策" class="toc-anchor">¶</a> 转诊和医疗保险政策</h2>
<h3 class="toc-header" id="转诊"><a href="#转诊" class="toc-anchor">¶</a> 转诊</h3>
<p>若校医院无法满足你的医疗需求/你认为需要前往更高级别医院进行治疗，且你享有医保政策时，你可能需要进行转诊服务。<br>
如不需要使用医保政策，不需要进行转诊。<br>
如当年（自然年）已达到300元门诊起付标准，不需要进行转诊。</p>
<h4 class="toc-header" id="校医院正常服务时"><a href="#校医院正常服务时" class="toc-anchor">¶</a> 校医院正常服务时</h4>
<p><strong>一、在校医院</strong><br>
通过<strong>医保挂号</strong>后，向接诊医生说明病情和想转诊的医院。转诊后，门诊医保起付标准减免300元。</p>
<p><strong>二、在所转诊的医院</strong><br>
在转诊有效日期内，转诊到不同等级医院就诊时，使用<strong>医保卡/市民卡/电子医保凭证</strong>结算，才能享受大学生医保待遇，享受在医保范围内<strong>门诊</strong>，三级医院，医保基金承担40%，其他等级医疗机构医保基金承担60%，社区卫生服务中心医保基金承担70%。如对费用有疑问，请现场咨询医院收费人员或医保办人员，现场解决。</p>
<p><strong>三、多次就诊</strong><br>
在转诊有效日期内，并在统一转诊医院内，可多科室，多次就诊（无需回校重复转诊）。</p>
<h4 class="toc-header" id="校医院不在服务时间或暂停服务时"><a href="#校医院不在服务时间或暂停服务时" class="toc-anchor">¶</a> 校医院不在服务时间或暂停服务时</h4>
<p>先使用<strong>医保卡/市民卡/电子医保凭证</strong>前往校外医保定点医药机构进行就诊。就诊结束支付时，将根据医保政策自动结算，支付金额即为个人承担部分金额。</p>
<h3 class="toc-header" id="医疗保险政策"><a href="#医疗保险政策" class="toc-anchor">¶</a> 医疗保险政策</h3>
<p>以下内容较为复杂，由于关系到同学们的切身利益，请务必仔细阅读。</p>
<h4 class="toc-header" id="什么是医保"><a href="#什么是医保" class="toc-anchor">¶</a> 什么是医保？</h4>
<p>我们通常说的医保，是指<strong>城乡居民基本医疗保险</strong>，不同城市、不同省、不同人群的参保费用、报销占比、结算（特别是异地结算）规则<strong>非常复杂</strong>！因此，本文的“医保”报销相关政策，均<strong>特指</strong>符合绝大多数同学的参保政策：<strong>杭州市基本医疗保障办法-城乡居民医保-大学生医保</strong>。如果你的参保情况不属于该情况，请咨询校医院和<strong>属地医保机构</strong>。</p>
<p>如开头所述，入学后，学校会统一安排以班级为单位缴纳保险费用，包含两部分（杭州市城乡居民医保-大学生医保、补充商业保险）。这个价格非常便宜（医保执行标准：每人每年270元，其中个人缴纳90元，同级财政补贴180元，一次性四年360元。补充商业保险：30元一年，一次性四年120元。共计480元。），<strong>一定要办！</strong>  这个保险是用来<strong>兜底</strong>的，对于经济困难的同学更加需要，确实存在经济困难的，请<strong>向学校寻求帮助！</strong>（拿一些常见的门诊检测项目举例，血液类检查（单项）、X平扫在几十到一两百，超声在200-300，CT、MRI一般在500-数千）</p>
<p>杭州市城乡居民医保-大学生医保 购买说明：</p>
<ol>
<li>持有效期内杭州市《特困人员救助供养证》、《最低生活保障家庭证》、《残疾人基本生活保障证》和二级及以上《中华人民共和国残疾人证》的学生，免缴当年度医保费；持有效期内杭州市《最低生活保障边缘家庭证》的学生，只需缴纳当年度一半保费，即45元。符合条件的学生请在规定缴费时间内携带相关证件向校医保办提交申请。</li>
<li>在9月新生入学，办理大学生新参时，省内原参保地有有效参保状态的城乡居民医保时，如何办理新参保手续? 答:根据全省统一城乡居民参(续)保管理业务规则，在集中征缴期办理本年度城乡居民参保登记时，省内存在其他地区城乡居民有效参保状态，本年度不予办理城乡居民参保登记，包括年度内大学生医保。</li>
<li>在9月新生入学，办理大学生新参时，省外原参保地有有效参保状态的城乡居民医保时，如何办理新参保手续? 答:根据全省统一城乡居民参(续)保管理业务规则，在集中征缴期办理本年度城乡居民参保登记时，省外存在其他地区城乡居民有效参保状态，原则上不建议新参杭州大学生医保。如遇省外医保在杭确实无法结算享受待遇的情况下，在办理原参保地医保停保后，依规申请参加杭州大学生医保。</li>
</ol>
<p>更多问答，参见<a class="is-external-link" href="https://xyy.hdu.edu.cn/2023/0612/c506a247374/page.htm">校医院2023年大学生医保问答</a>。</p>
<h4 class="toc-header" id="杭州市城乡居民医保-大学生医保"><a href="#杭州市城乡居民医保-大学生医保" class="toc-anchor">¶</a> 杭州市城乡居民医保-大学生医保</h4>
<p>任何与本文书写时现行《杭州市基本医疗保障办法 杭政〔2020〕56号》描述不一致的，以《办法》为准。<br>
若有更新（本版废止）的情况，以更新版本为准。<br>
《办法》全文链接：<br>
<a class="is-external-link" href="http://www.hangzhou.gov.cn/art/2020/12/31/art_1229063381_1716370.html">杭州市人民政府 《杭州市人民政府关于印发杭州市基本医疗保障办法的通知》</a><br>
<a class="is-external-link" href="https://www.zj.gov.cn/zjservice/item/detail/lawtext.do?outLawId=269dff1a-d0d8-41fb-8a0a-5279f400a87b">浙江省政务服务网（浙里办） 转发版本</a></p>
<h5 class="toc-header" id="重点结算规则"><a href="#重点结算规则" class="toc-anchor">¶</a> 重点：结算规则</h5>
<p>（一）门诊医疗费用结算</p>
<p>1、先由个人承担一个门诊起付标准300元。</p>
<p>2、门诊起付标准以上部分医疗费用由统筹基金和个人共同承担，具体如下：</p>
<div class="table-container"><table>
<thead>
<tr>
<th>比例</th>
<th>三级医疗机构</th>
<th>其他医疗机构（含二级）</th>
<th>社区卫生服务机构</th>
</tr>
</thead>
<tbody>
<tr>
<td>统筹基金比例</td>
<td>40%</td>
<td>60%</td>
<td>70%</td>
</tr>
<tr>
<td>个人承担比例</td>
<td>60%</td>
<td>40%</td>
<td>30%</td>
</tr>
</tbody>
</table></div>
<p>3、大学生在本校校医院门诊就医时，统筹基金承担比例在规定基础上提高3%，即个人承担比例为27%。</p>
<p>4、大学生在校医院首诊，或经校医院转诊至其他医疗机构继续治疗，门诊医保起付标准减免300元。</p>
<p>（二）住院医疗费用结算</p>
<p>1、由个人承担一个住院起付标准，两次及以上住院的，起付标准按其中最高等级医疗机构标准计算。三级医疗机构800元，其他医疗机构500元，社区卫生服务机构300元。</p>
<p>2、住院起付标准以上部分医疗费，由统筹基金和个人共同承担。承担比例如下：</p>
<div class="table-container"><table>
<thead>
<tr>
<th>比例</th>
<th>三级医疗机构</th>
<th>其他医疗机构（含二级）</th>
<th>社区卫生服务机构</th>
</tr>
</thead>
<tbody>
<tr>
<td>统筹基金比例</td>
<td>70%</td>
<td>75%</td>
<td>80%</td>
</tr>
<tr>
<td>个人承担比例</td>
<td>30%</td>
<td>25%</td>
<td>20%</td>
</tr>
</tbody>
</table></div>
<h5 class="toc-header" id="其它规则"><a href="#其它规则" class="toc-anchor">¶</a> 其它规则</h5>
<p>大学生医保其它相关规则，例如住院报销等，请<strong>参考</strong>校医院<a class="is-external-link" href="https://xyy.hdu.edu.cn/2021/0908/c506a136753/page.htm">杭州市大学生医保相关规定</a>，其中任何与本文书写时现行《杭州市基本医疗保障办法 杭政〔2020〕56号》描述不一致的，以《办法》为准。若有更新（本版废止）的情况，以更新版本为准。</p>
<h4 class="toc-header" id="通过学校购买的补充商业保险"><a href="#通过学校购买的补充商业保险" class="toc-anchor">¶</a> 通过学校购买的补充商业保险</h4>
<p>保单上会说明相关保障项目和报销比例，主要覆盖意外伤害和住院医疗（以21年所售为例，在同时购买医保情况下，住院医疗医保结算后个人部分，在最高限额内全额赔付。以实际签署的保险合同为准。），性价比相对市售保险高很多。<br>
根据校医院说明，<strong>通过学校购买的补充商业保险</strong>，2018级-2020级学生需线下办理，参见<a class="is-external-link" href="https://xyy.hdu.edu.cn/2017/0810/c506a31505/page.htm">2018级-2020级学生申请商业补充保险报销的准备材料</a>;2021级-2022级学生通过小程序“平安保险好生活”或APP线上理赔即可；后续学生理赔方式，若商保公司更换，可能更换理赔方式，以校医院或保险合同为准。<br>
医保办的上班时间是：周一到周五 8:00-11:30 13:30-16:00</p>
<h4 class="toc-header" id="商业保险"><a href="#商业保险" class="toc-anchor">¶</a> 商业保险</h4>
<p>商业保险规则非常复杂。由于医保相关政策变化（参见CHS-DRG疾病诊断相关分组分组与付费技术规范政策，细节不便展开），为满足医疗需求，很多地区出现了“政府指导的地区限定商业医疗保险”（例如杭州市-西湖益联保），这些保险通常合同相对更可靠、与当地医院医保结算系统相对更紧密、价格较普惠，请注意，本句话仅做情况说明，不意味着任何推荐、营销或潜在营销行为，具体信息请参考目标地有关部门公开说明，自行决定是否购买。同时，<strong>请在购买任何保险服务前，认真阅读和理解保险合同及连带文件，避免自身利益受损</strong>。</p>
<p><strong>相关链接</strong></p>
<p>更多内容，请访问<a class="is-external-link" href="https://xyy.hdu.edu.cn/506/list.htm">校医院学生医保政策说明专区</a>获取更多信息。</p>
`,
			},
		})
	})
}
