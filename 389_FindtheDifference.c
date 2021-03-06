/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了52.56%的用户
*内存消耗：5.6 MB, 在所有 Go 提交中击败了38.22%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。
************************************************************************************************************
*/


char findTheDifference(char* s, char* t) {
    int len = 0;
    int i;
    char *ptr = NULL;
    char tmp = 0;
    
    len = strlen(t);
    char *buff = (char *)malloc(sizeof(char)*(len+1));
    memset(buff, 0, sizeof(char)*(len+1));
    memcpy(buff, t, len);
    for(i = 0; i < len-1; i++){
        if((ptr = strchr(buff, s[i])) != NULL){
            *ptr = ' ';
        }
    }
    
    for(i = 0; i < len; i++){
        if(buff[i] != ' '){
            tmp = buff[i];
            free(buff);
            break;
        }
    }
    return tmp;
}