namespace QoraUI
{
    partial class FormQoraUI
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.tabControlForm = new System.Windows.Forms.TabControl();
            this.tabPageEncrypt = new System.Windows.Forms.TabPage();
            this.tabPageDecrypt = new System.Windows.Forms.TabPage();
            this.tabPageConfigure = new System.Windows.Forms.TabPage();
            this.tabPageAbout = new System.Windows.Forms.TabPage();
            this.tabControlForm.SuspendLayout();
            this.SuspendLayout();
            // 
            // tabControlForm
            // 
            this.tabControlForm.Controls.Add(this.tabPageEncrypt);
            this.tabControlForm.Controls.Add(this.tabPageDecrypt);
            this.tabControlForm.Controls.Add(this.tabPageConfigure);
            this.tabControlForm.Controls.Add(this.tabPageAbout);
            this.tabControlForm.Location = new System.Drawing.Point(12, 12);
            this.tabControlForm.Name = "tabControlForm";
            this.tabControlForm.SelectedIndex = 0;
            this.tabControlForm.Size = new System.Drawing.Size(750, 505);
            this.tabControlForm.TabIndex = 0;
            // 
            // tabPageEncrypt
            // 
            this.tabPageEncrypt.Location = new System.Drawing.Point(8, 39);
            this.tabPageEncrypt.Name = "tabPageEncrypt";
            this.tabPageEncrypt.Padding = new System.Windows.Forms.Padding(3);
            this.tabPageEncrypt.Size = new System.Drawing.Size(734, 458);
            this.tabPageEncrypt.TabIndex = 0;
            this.tabPageEncrypt.Text = "加密";
            this.tabPageEncrypt.UseVisualStyleBackColor = true;
            // 
            // tabPageDecrypt
            // 
            this.tabPageDecrypt.Location = new System.Drawing.Point(8, 39);
            this.tabPageDecrypt.Name = "tabPageDecrypt";
            this.tabPageDecrypt.Padding = new System.Windows.Forms.Padding(3);
            this.tabPageDecrypt.Size = new System.Drawing.Size(734, 458);
            this.tabPageDecrypt.TabIndex = 1;
            this.tabPageDecrypt.Text = "解密";
            this.tabPageDecrypt.UseVisualStyleBackColor = true;
            // 
            // tabPageConfigure
            // 
            this.tabPageConfigure.Location = new System.Drawing.Point(8, 39);
            this.tabPageConfigure.Name = "tabPageConfigure";
            this.tabPageConfigure.Padding = new System.Windows.Forms.Padding(3);
            this.tabPageConfigure.Size = new System.Drawing.Size(734, 458);
            this.tabPageConfigure.TabIndex = 2;
            this.tabPageConfigure.Text = "设置";
            this.tabPageConfigure.UseVisualStyleBackColor = true;
            // 
            // tabPageAbout
            // 
            this.tabPageAbout.Location = new System.Drawing.Point(8, 39);
            this.tabPageAbout.Name = "tabPageAbout";
            this.tabPageAbout.Padding = new System.Windows.Forms.Padding(3);
            this.tabPageAbout.Size = new System.Drawing.Size(734, 458);
            this.tabPageAbout.TabIndex = 3;
            this.tabPageAbout.Text = "关于";
            this.tabPageAbout.UseVisualStyleBackColor = true;
            // 
            // FormQoraUI
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(12F, 24F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(774, 529);
            this.Controls.Add(this.tabControlForm);
            this.MaximizeBox = false;
            this.Name = "FormQoraUI";
            this.Text = "QoraUI";
            this.Load += new System.EventHandler(this.FormQoraUI_Load);
            this.tabControlForm.ResumeLayout(false);
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TabControl tabControlForm;
        private System.Windows.Forms.TabPage tabPageEncrypt;
        private System.Windows.Forms.TabPage tabPageDecrypt;
        private System.Windows.Forms.TabPage tabPageConfigure;
        private System.Windows.Forms.TabPage tabPageAbout;
    }
}

