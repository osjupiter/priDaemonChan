$methods = @'
[return: MarshalAs(UnmanagedType.Bool)]
[DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
public static extern bool PostMessage(uint hWnd, uint Msg, IntPtr wParam, IntPtr lParam);

[DllImport("user32.dll", SetLastError = true)]
public static extern bool LockWorkStation();
'@

$pwm = Add-Type -MemberDefinition $methods -Name "PowerManager" -PassThru -Language CSharp

Start-Sleep -m 3000

#$pwm::PostMessage(0xffff, 0x0112, 0xf170, 2)
#$pwm::LockWorkStation()
$pwm::PostMessage(0xffff, 0x0112, 0xf170, -1)