bool isPattern(char *buff, int len, int count){

    int i;
    char *ptr = buff;
    //printf("count=%d\n", count);
    for(i = count; i < len; i++){
        //printf("%d %d %c %c\n", i%count, i, buff[i%count], ptr[i]);
        if(buff[i%count] != *(ptr+i)){
            return false;
        }
    }

    return true;
}

bool repeatedSubstringPattern(char * s){

    if (s == NULL | strlen(s) <= 0) {
        return false;
    }

    int i;
    int len;
    int count = 0;
    len = strlen(s);

    for(i = 1; i <= len/2; i++){

        if(len%i != 0){
            continue;
        }

        if(isPattern(s, len, i) == true){
            return true;
        }

    }

    return false;

}