@echo off

rem ��������Ϊ �������� ��ִ���ļ�����
@echo ��������:%1
@echo ��������:%2

rem ������Ҫ���еĳ���·��
set curExe=%~dp0%2
rem ����ע���·��
set regpath=HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%1\Parameters\
rem ����srvany.exe�ļ�·��
set sourcePath=%~dp0tasksch.exe

rem ���뵱ǰĿ¼
cd /d "%~dp0"
rem ��װ��������
instsrv %1  "%sourcePath%"
@echo ����������

rem ���ע����﷨: reg add ע���·�� /v ������ /t ֵ���� /d ���� /f ��ʾǿ���޸Ĳ���ʾ

rem ���� Application ֵΪ��Ҫ��Ϊ�������еĳ����ַ /d��Ӧ�Ĳ�����б�ܲ���Ϊ��ת�����ţ�����·������б�ܣ�Ĭ�Ͻ�����ת���ˣ��������б����Ϊ�˱�������
reg add %regpath% /v AppDirectory /t REG_SZ /d "%~dp0\" /f

rem ���� AppDirectory ֵΪ��Ҫ��Ϊ�������еĳ��������ļ���·��
reg add %regpath% /v Application /t REG_SZ /d "%curExe%" /f 

rem ���� AppParameters ֵΪ��Ҫ��Ϊ�������еĳ�����������Ҫ�Ĳ���
reg add %regpath% /v AppParameters /t REG_SZ /f
@echo ע���������