Vim�UnDo� ��D��'%�Z!UKl_G��;R��8�
Y҃ڕ(Z   '           "      *       *   *   *    _B�2    _�                             ����                                                                                                                                                                                                                                                                                                                                                             _B�[     �                  �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _B�n     �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B     �             5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             _B     �               func Run() string {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B      �                	tempCommand := getFunctor(Lang)5�_�                       $    ����                                                                                                                                                                                                                                                                                                                                                             _B®     �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _B³     �                 5�_�      	                      ����                                                                                                                                                                                                                                                                                                                                                             _B³     �               	�             5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                             _B¶    �                  package program       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }            func Run(prog *Program) string {   %	tempCommand := getFunctor(prog.Lang)       	return "hi"   }5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                             _B»     �               	�             5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             _B��     �               %	tempCommand := getFunctor(prog.Lang)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B��     �                	prog5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _B��     �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B��     �               	runerFileFunctor(5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B��     �               	runerFileFunctor(prog)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _B��    �                  package program       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }        func Run(prog *Program) string {   *	runerFileFunctor := getFunctor(prog.Lang)       &	tempCommand := runerFileFunctor(prog)   	return "hi"   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B��     �               *	runerFileFunctor := getFunctor(prog.Lang)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B��     �               &	tempCommand := runerFileFunctor(prog)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B��    �                  package program       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }        func Run(prog *Program) string {   +	runnerFileFunctor := getFunctor(prog.Lang)       '	tempCommand := runnerFileFunctor(prog)   	return "hi"   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B��     �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B�&     �             �               	�             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B�I     �         !      	new�                	newCommand.WriteString5�_�                           ����                                                                                                                                                                                                                                                                                                                                                v       _B�k     �                
	new(type)5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        _B�n     �              �              5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        _B�q     �         !      $	newCommand.WriteString(tempCommand)5�_�                            ����                                                                                                                                                                                                                                                                                                                                                V       _B�v     �         !    �         !    5�_�                           ����                                                                                                                                                                                                                                                                                                                                                V       _B�y     �         "      	newCommand.WriteString(">> ")5�_�                            ����                                                                                                                                                                                                                                                                                                                               #          #       V   &    _BÈ     �          "    �         "    5�_�                           ����                                                                                                                                                                                                                                                                                                                               #          #       V   &    _BË     �          #      &	newCommand.WriteString("output.txt ")5�_�                    !        ����                                                                                                                                                                                                                                                                                                                               #          #       V   &    _BÑ    �               #   package program       import "strings"       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }        func Run(prog *Program) string {   +	runnerFileFunctor := getFunctor(prog.Lang)       '	tempCommand := runnerFileFunctor(prog)       	var newCommand strings.Builder       $	newCommand.WriteString(tempCommand)   	newCommand.WriteString(">> ")   &	newCommand.WriteString("output.txt ")   	newCommand.WriteString("2>&1")           	return "hi"   }5�_�                            ����                                                                                                                                                                                                                                                                                                                               #          #       V   &    _BÙ     �      "   #      	�      !   "    5�_�      !               !        ����                                                                                                                                                                                                                                                                                                                               #          #       V   &    _Bï     �         %      import "strings"�   !   #   %      	�   !   #   $    5�_�       "           !   %       ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _B�     �   $   &   (      	exec(5�_�   !   #           "   %       ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _B�"     �   $   &   (      	out, err := exec(5�_�   "   $           #   %   +    ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _B�E     �   $   &   (      ,	out, err := exec.Command(newCommand.String)5�_�   #   %           $   %   .    ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _B�i     �   $   &   (      .	out, err := exec.Command(newCommand.String())5�_�   $   &           %          ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _BČ     �         (      '	tempCommand := runnerFileFunctor(prog)5�_�   %   '           &      "    ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _BĒ     �         (      "	tempCommand := runnerFileFunctor(5�_�   &   (           '      "    ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _Bė     �         (      "	tempCommand := runnerFileFunctor(5�_�   '   )           (   $        ����                                                                                                                                                                                                                                                                                                                            !   #       !   #       V   &    _BĚ    �               (   package program       import (   
	"os/exec"   
	"strings"   )       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }        func Run(prog *Program) string {   +	runnerFileFunctor := getFunctor(prog.Lang)       (	tempCommand := runnerFileFunctor(*prog)       	var newCommand strings.Builder       $	newCommand.WriteString(tempCommand)   	newCommand.WriteString(">> ")   &	newCommand.WriteString("output.txt ")   	newCommand.WriteString("2>&1")           7	out, err := exec.Command(newCommand.String()).Output()       	return "hi"   }5�_�   (   *           )      "    ����                                                                                                                                                                                                                                                                                                                                                             _B�0     �         '      (	tempCommand := runnerFileFunctor(*prog)5�_�   )               *      !    ����                                                                                                                                                                                                                                                                                                                                                             _B�1    �               '   package program       import (   
	"os/exec"   
	"strings"   )       type Program struct {   	Lang       string   	Code       string   	FileRunner string   	OutputFile string   }       3func NewProgram(lang string, code string) Program {   	prog := Program{   		Lang: lang,   		Code: code,   	}       	return prog   }        func Run(prog *Program) string {   +	runnerFileFunctor := getFunctor(prog.Lang)       #	tempCommand := runnerFileFunctor()       	var newCommand strings.Builder       $	newCommand.WriteString(tempCommand)   	newCommand.WriteString(">> ")   &	newCommand.WriteString("output.txt ")   	newCommand.WriteString("2>&1")       7	out, err := exec.Command(newCommand.String()).Output()       	return "hi"   }5��